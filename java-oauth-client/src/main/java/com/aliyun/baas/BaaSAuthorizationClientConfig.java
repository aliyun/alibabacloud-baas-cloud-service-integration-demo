package com.aliyun.baas;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.HttpHeaders;
import org.springframework.http.client.ClientHttpRequestFactory;
import org.springframework.http.client.SimpleClientHttpRequestFactory;
import org.springframework.security.access.AccessDeniedException;
import org.springframework.security.oauth2.client.DefaultOAuth2ClientContext;
import org.springframework.security.oauth2.client.OAuth2RestOperations;
import org.springframework.security.oauth2.client.OAuth2RestTemplate;
import org.springframework.security.oauth2.client.resource.BaseOAuth2ProtectedResourceDetails;
import org.springframework.security.oauth2.client.resource.OAuth2AccessDeniedException;
import org.springframework.security.oauth2.client.resource.OAuth2ProtectedResourceDetails;
import org.springframework.security.oauth2.client.resource.UserRedirectRequiredException;
import org.springframework.security.oauth2.client.token.AccessTokenProvider;
import org.springframework.security.oauth2.client.token.AccessTokenRequest;
import org.springframework.security.oauth2.client.token.DefaultAccessTokenRequest;
import org.springframework.security.oauth2.client.token.OAuth2AccessTokenSupport;
import org.springframework.security.oauth2.common.DefaultOAuth2AccessToken;
import org.springframework.security.oauth2.common.DefaultOAuth2RefreshToken;
import org.springframework.security.oauth2.common.OAuth2AccessToken;
import org.springframework.security.oauth2.common.OAuth2RefreshToken;
import org.springframework.util.LinkedMultiValueMap;
import org.springframework.util.MultiValueMap;

@Configuration
public class BaaSAuthorizationClientConfig {
    @Autowired(required = false)
    ClientHttpRequestFactory clientHttpRequestFactory;

    /*
     * ClientHttpRequestFactory is autowired and checked in case somewhere in
     * your configuration you provided {@link ClientHttpRequestFactory}
     * implementation Bean where you defined specifics of your connection, if
     * not it is instantiated here with {@link SimpleClientHttpRequestFactory}
     */
    private ClientHttpRequestFactory getClientHttpRequestFactory() {
        if (clientHttpRequestFactory == null) {
            clientHttpRequestFactory = new SimpleClientHttpRequestFactory();
        }
        return clientHttpRequestFactory;
    }

    @Bean
    @Qualifier("baasRestTemplate")
    public OAuth2RestOperations restTemplate(@Value("${baas.rest.endpoint}") String restUrl, @Value("${baas.rest.refresh_token}") String refreshToken) {
        String tokenUrl = restUrl + "/api/v1/token";
        OAuth2RestTemplate template = new OAuth2RestTemplate(refreshTokenResourceDetails(tokenUrl, refreshToken), new DefaultOAuth2ClientContext(new DefaultAccessTokenRequest()));
        template.setRequestFactory(getClientHttpRequestFactory());
        template.setAccessTokenProvider(new RefreshTokenAccessTokenProvider());
        return template;
    }

    public OAuth2ProtectedResourceDetails refreshTokenResourceDetails(String tokenUrl, String refreshToken) {
        RefreshTokenResourceDetails resource = new RefreshTokenResourceDetails();
        resource.setAccessTokenUri(tokenUrl);
        resource.setRefreshToken(refreshToken);
        return resource;
    }

    public class RefreshTokenAccessTokenProvider extends OAuth2AccessTokenSupport implements AccessTokenProvider {
        public RefreshTokenAccessTokenProvider() {
        }

        public boolean supportsResource(OAuth2ProtectedResourceDetails resource) {
            return resource instanceof RefreshTokenAccessTokenProvider && "refresh_token".equals(resource.getGrantType());
        }

        public boolean supportsRefresh(OAuth2ProtectedResourceDetails resource) {
            return this.supportsResource(resource);
        }

        public OAuth2AccessToken refreshAccessToken(OAuth2ProtectedResourceDetails resource, OAuth2RefreshToken refreshToken, AccessTokenRequest request) throws UserRedirectRequiredException, OAuth2AccessDeniedException {
            MultiValueMap<String, String> form = new LinkedMultiValueMap();
            form.add("grant_type", "refresh_token");
            form.add("refresh_token", refreshToken.getValue());
            return this.retrieveToken(request, resource, form, new HttpHeaders());
        }

        public OAuth2AccessToken obtainAccessToken(OAuth2ProtectedResourceDetails details, AccessTokenRequest request) throws UserRedirectRequiredException, AccessDeniedException, OAuth2AccessDeniedException {
            RefreshTokenResourceDetails resource = (RefreshTokenResourceDetails)details;
            DefaultOAuth2RefreshToken refreshToken = new DefaultOAuth2RefreshToken(resource.getRefreshToken());
            OAuth2AccessToken newToken = this.refreshAccessToken(resource, refreshToken, request);
            DefaultOAuth2AccessToken token = new DefaultOAuth2AccessToken(newToken.getValue());
            token.setRefreshToken(refreshToken);
            token.setAdditionalInformation(newToken.getAdditionalInformation());
            token.setExpiration(newToken.getExpiration());
            token.setScope(newToken.getScope());
            token.setTokenType(newToken.getTokenType());
            return token;
        }
    }

    public class RefreshTokenResourceDetails extends BaseOAuth2ProtectedResourceDetails {
        private String refreshToken;

        public RefreshTokenResourceDetails() {
            this.setGrantType("refresh_token");
        }

        public String getRefreshToken() {
            return refreshToken;
        }

        public void setRefreshToken(String refreshToken) {
            this.refreshToken = refreshToken;
        }
    }
}
