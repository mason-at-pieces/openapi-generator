# NOTE: This file is auto generated by OpenAPI Generator 7.5.0-SNAPSHOT (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule OpenapiPetstore.Model.HasOnlyReadOnly do
  @moduledoc """
  
  """

  @derive Jason.Encoder
  defstruct [
    :bar,
    :foo
  ]

  @type t :: %__MODULE__{
    :bar => String.t | nil,
    :foo => String.t | nil
  }

  def decode(value) do
    value
  end
end

