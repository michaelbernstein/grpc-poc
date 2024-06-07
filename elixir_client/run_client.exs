# run_client.exs

Mix.Task.run("app.start")

defmodule RunClient do
  def run do
    # Call the query function
    IO.puts("Running single query...")
    PanopticonClient.query()
  end
end

RunClient.run()

