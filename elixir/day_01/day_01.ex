defmodule Day01 do
  defp read_input(input) do
    Path.join(["..", "..", "inputs", "day_01", input])
    |> File.read!()
    |> String.replace("\r\n", "\n")
    |> String.split("\n")
    |> Enum.filter(&(&1 != ""))
  end

  defp first_and_last(list) do
    [Enum.at(list, 0), Enum.at(list, -1)]
  end

  defp process_line(line) do
    line
    # Get each character
    |> String.graphemes()
    # Filter out non-numbers
    |> Enum.filter(fn x -> String.match?(x, ~r/^-?\d+$/) end)
    # Convert to integers
    |> Enum.map(&String.to_integer/1)
    |> first_and_last()
    # Convert to a single number
    |> Enum.reduce(&(&1 + &2 * 10))
  end

  def part_one do
    read_input("input.txt")
    |> Enum.map(&process_line/1)
    |> Enum.sum()
  end

  defp process_line_two(line) do
    line
    |> String.replace("one", "o1e")
    |> String.replace("two", "t2o")
    |> String.replace("three", "t3e")
    |> String.replace("four", "f4r")
    |> String.replace("five", "f5e")
    |> String.replace("six", "s6x")
    |> String.replace("seven", "s7n")
    |> String.replace("eight", "e8t")
    |> String.replace("nine", "n9e")
    |> process_line()
    |> IO.inspect()
  end

  def part_two do
    read_input("input.txt")
    |> Enum.map(&process_line_two/1)
    |> Enum.sum()
  end
end

Day01.part_two() |> IO.inspect()
