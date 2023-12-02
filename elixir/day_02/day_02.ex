defmodule Day02 do
  defp read_input(input) do
    Path.join(["..", "..", "inputs", "day_02", input])
    |> File.read!()
    |> String.replace("\r\n", "\n")
    |> String.split("\n", trim: true)
  end

  defp validate_reveal(count, color) do
    case color do
      "red" -> count <= 12
      "green" -> count <= 13
      "blue" -> count <= 14
      _ -> false
    end
  end

  defp parse_reveal(reveal) do
    reveal
    |> String.split(",", trim: true)
    |> Enum.map(&String.trim/1)
    |> Enum.map(&String.split(&1, " ", trim: true))
    |> Enum.map(fn [count, color] ->
      count = String.to_integer(count)
      {count, color}
    end)
  end

  defp parse_game(game) do
    game
    |> String.split(";", trim: true)
    |> Enum.map(&String.trim/1)
    |> Enum.map(&parse_reveal/1)
  end

  defp process_line(line) do
    [game_id, results] =
      line
      |> String.split(":", trim: true)
      |> Enum.map(&String.trim/1)

    game_id =
      game_id
      |> String.split(" ", trim: true)
      |> Enum.at(1)
      |> String.to_integer()

    game = results |> parse_game()

    [game_id, game]
  end

  def part_one do
    read_input("input.txt")
    |> Enum.map(&process_line/1)
    |> Enum.map(fn [game_id, game] ->
      [
        game_id,
        Enum.all?(game, fn reveals ->
          reveals
          |> Enum.map(fn {count, color} -> validate_reveal(count, color) end)
          |> Enum.all?()
        end)
      ]
    end)
    |> Enum.filter(fn [_, results] -> results end)
    |> Enum.map(fn [game_id, _] -> game_id end)
    |> Enum.sum()
  end

  def part_two do
    read_input("input.txt")
    |> Enum.map(&process_line/1)
    |> Enum.map(fn [_, game] ->
      game
      |> Enum.reduce(
        %{
          "red" => 0,
          "green" => 0,
          "blue" => 0
        },
        fn reveals, acc ->
          reveals
          |> Enum.reduce(acc, fn {count, color}, acc ->
            Map.update(acc, color, count, &max(&1, count))
          end)
        end
      )
      |> Map.values()
      |> Enum.product()
    end)
    |> Enum.sum()
  end
end

Day02.part_one() |> IO.inspect()
Day02.part_two() |> IO.inspect()
