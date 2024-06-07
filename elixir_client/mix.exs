defmodule ElixirClient.MixProject do
  use Mix.Project

  def project do
    [
      app: :elixir_client,
      version: "0.1.0",
      elixir: "~> 1.11",
      start_permanent: Mix.env() == :prod,
      deps: deps()
    ]
  end

  def application do
    [
      extra_applications: [:logger]
    ]
  end

  defp deps do
    [
      {:grpc, github: "elixir-grpc/grpc"},
      {:protobuf, "~> 0.11"},
      {:google_protos, "~> 0.2"}
    ]
  end
end

