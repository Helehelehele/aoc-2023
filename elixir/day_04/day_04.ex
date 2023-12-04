defmodule Day04 do
  defp read_input(input) do
    Path.join(["..", "..", "inputs", "day_04", input])
    |> File.read!()
    |> String.replace("\r\n", "\n")
    |> String.split("\n", trim: true)
  end

  defp convert_to_integer_list(string_list) do
    string_list
    |> Enum.map(&String.to_integer/1)
  end

  defp convert_to_sets(list_of_lists) do
    list_of_lists
    |> Enum.map(&MapSet.new/1)
  end

  defp intersect_sets(list_of_sets) do
    list_of_sets
    |> Enum.reduce(&MapSet.intersection(&1, &2))
  end

  defp process_line({line, index}) do
    line
    |> String.split(": ", trim: true)
    |> Enum.at(1)
    |> String.split(" | ", trim: true)
    |> Enum.map(&String.split(&1, " ", trim: true))
    |> Enum.map(&convert_to_integer_list/1)
    |> then(fn list -> {index, list} end)
  end

  defp win_copies(count_map, index_seq, total_cards) do
    case {index_seq, total_cards} do
      {[], ^total_cards} ->
        total_cards

      {[head | tail], ^total_cards} ->
        0..Map.get(count_map, head, 0)
        |> Enum.filter(&(&1 > 0))
        |> Enum.map(fn x -> head + x end)
        |> Enum.filter(fn x -> Map.has_key?(count_map, x) end)
        |> Enum.reverse()
        |> Enum.reduce(tail, fn x, acc ->
          [x | acc]
        end)
        |> then(fn acc ->
          {count_map, acc, total_cards + 1}
        end)
        |> then(fn {count_map, acc, total_cards} ->
          win_copies(count_map, acc, total_cards)
        end)
    end
  end

  def part_one do
    read_input("input.txt")
    |> Enum.with_index(1)
    |> Enum.map(&process_line/1)
    |> Enum.map(&elem(&1, 1))
    |> Enum.map(&convert_to_sets/1)
    |> Enum.map(&intersect_sets/1)
    |> Enum.map(&MapSet.size/1)
    |> Enum.filter(&(&1 > 0))
    |> Enum.map(fn size -> max(0, size - 1) end)
    |> Enum.map(fn size -> :math.pow(2, size) end)
    |> Enum.sum()
    |> trunc()
  end

  def part_two do
    read_input("input.txt")
    |> Enum.with_index(1)
    |> Enum.map(&process_line/1)
    |> Enum.map(fn {index, list} -> {index, convert_to_sets(list)} end)
    |> Enum.map(fn {index, sets} -> {index, intersect_sets(sets)} end)
    |> Enum.map(fn {index, intersections} -> {index, MapSet.size(intersections)} end)
    # Turn this into a map
    |> Enum.reduce(%{}, fn {index, size}, acc -> Map.put(acc, index, size) end)
    |> then(fn count_map ->
      win_copies(count_map, Map.keys(count_map) |> Enum.sort(), 0)
    end)
  end
end

Day04.part_one() |> IO.puts()
Day04.part_two() |> IO.puts()
