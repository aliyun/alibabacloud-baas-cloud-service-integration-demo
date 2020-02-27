/*
 * BaaS API Gateway
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */

package com.aliyun.baas.rest.module;

import java.util.Objects;
import java.util.Arrays;
import com.aliyun.baas.rest.module.Error;
import com.aliyun.baas.rest.module.InlineResponse2003Result;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;
import io.swagger.v3.oas.annotations.media.Schema;
/**
 * InlineResponse2003
 */

@javax.annotation.Generated(value = "io.swagger.codegen.v3.generators.java.JavaClientCodegen", date = "2020-02-12T03:32:15.669Z[GMT]")
public class InlineResponse2003 {
  @JsonProperty("Success")
  private Boolean success = null;

  @JsonProperty("Result")
  private InlineResponse2003Result result = null;

  @JsonProperty("Error")
  private Error error = null;

  public InlineResponse2003 success(Boolean success) {
    this.success = success;
    return this;
  }

   /**
   * Boolean indicating if the request was successful.
   * @return success
  **/
  @Schema(required = true, description = "Boolean indicating if the request was successful.")
  public Boolean isSuccess() {
    return success;
  }

  public void setSuccess(Boolean success) {
    this.success = success;
  }

  public InlineResponse2003 result(InlineResponse2003Result result) {
    this.result = result;
    return this;
  }

   /**
   * Get result
   * @return result
  **/
  @Schema(required = true, description = "")
  public InlineResponse2003Result getResult() {
    return result;
  }

  public void setResult(InlineResponse2003Result result) {
    this.result = result;
  }

  public InlineResponse2003 error(Error error) {
    this.error = error;
    return this;
  }

   /**
   * Get error
   * @return error
  **/
  @Schema(required = true, description = "")
  public Error getError() {
    return error;
  }

  public void setError(Error error) {
    this.error = error;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineResponse2003 inlineResponse2003 = (InlineResponse2003) o;
    return Objects.equals(this.success, inlineResponse2003.success) &&
        Objects.equals(this.result, inlineResponse2003.result) &&
        Objects.equals(this.error, inlineResponse2003.error);
  }

  @Override
  public int hashCode() {
    return Objects.hash(success, result, error);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineResponse2003 {\n");
    
    sb.append("    success: ").append(toIndentedString(success)).append("\n");
    sb.append("    result: ").append(toIndentedString(result)).append("\n");
    sb.append("    error: ").append(toIndentedString(error)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}