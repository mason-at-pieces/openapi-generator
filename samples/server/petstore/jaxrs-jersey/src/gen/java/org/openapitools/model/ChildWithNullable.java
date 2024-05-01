/*
 * OpenAPI Petstore
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.model;

import java.util.Objects;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ParentWithNullable;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import javax.validation.constraints.*;
import javax.validation.Valid;

/**
 * ChildWithNullable
 */
@JsonPropertyOrder({
  ChildWithNullable.JSON_PROPERTY_OTHER_PROPERTY
})
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaJerseyServerCodegen", comments = "Generator version: 7.5.0-SNAPSHOT")
public class ChildWithNullable extends ParentWithNullable  {
  public static final String JSON_PROPERTY_OTHER_PROPERTY = "otherProperty";
  @JsonProperty(JSON_PROPERTY_OTHER_PROPERTY)
  private String otherProperty;

  public ChildWithNullable otherProperty(String otherProperty) {
    this.otherProperty = otherProperty;
    return this;
  }

  /**
   * Get otherProperty
   * @return otherProperty
   **/
  @JsonProperty(value = "otherProperty")
  @ApiModelProperty(value = "")
  
  public String getOtherProperty() {
    return otherProperty;
  }

  public void setOtherProperty(String otherProperty) {
    this.otherProperty = otherProperty;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ChildWithNullable childWithNullable = (ChildWithNullable) o;
    return super.equals(o) && Objects.equals(this.otherProperty, childWithNullable.otherProperty);
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), otherProperty);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ChildWithNullable {\n");
    sb.append("    ").append(toIndentedString(super.toString())).append("\n");
    sb.append("    otherProperty: ").append(toIndentedString(otherProperty)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }
}

