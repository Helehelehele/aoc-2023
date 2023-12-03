defmodule Day03 do
  defmodule Number do
    defstruct [:value, :row, :column]

    def get_start(number) do
      number.column
    end

    def get_end(number) do
      number.column + String.length(number.value) - 1
    end
  end

  defmodule Symbol do
    defstruct [:value, :row, :column]

    def is_adjacent?(symbol, number) do
      if abs(symbol.row - number.row) > 1 do
        false
      else
        if symbol.column >= Number.get_start(number) - 1 and
             symbol.column <= Number.get_end(number) + 1 do
          true
        else
          false
        end
      end
    end
  end

  defp read_input(input) do
    Path.join(["..", "..", "inputs", "day_03", input])
    |> File.read!()
    |> String.replace("\r\n", "\n")
    |> String.split("\n")
    |> Enum.filter(&(&1 != ""))
  end

  defp process_line({line, index}) do
    line
    |> String.graphemes()
    |> Enum.with_index()
    |> Enum.reduce(
      # numbers, symbols, number_in_progress
      {[], [], nil},
      fn {x, i}, {numbers, symbols, number_in_progress} ->
        case x in ["0", "1", "2", "3", "4", "5", "6", "7", "8", "9"] do
          true ->
            {
              numbers,
              symbols,
              if number_in_progress == nil do
                {x, i}
              else
                {
                  elem(number_in_progress, 0) <> x,
                  elem(number_in_progress, 1)
                }
              end
            }

          false ->
            case x do
              "." ->
                case number_in_progress do
                  nil ->
                    {numbers, symbols, number_in_progress}

                  {value, col} ->
                    {
                      [
                        %Number{
                          value: value,
                          row: index,
                          column: col
                        }
                        | numbers
                      ],
                      symbols,
                      nil
                    }
                end

              symbol ->
                case number_in_progress do
                  nil ->
                    {
                      numbers,
                      [
                        %Symbol{
                          value: symbol,
                          row: index,
                          column: i
                        }
                        | symbols
                      ],
                      number_in_progress
                    }

                  {value, col} ->
                    {
                      [
                        %Number{
                          value: value,
                          row: index,
                          column: col
                        }
                        | numbers
                      ],
                      [
                        %Symbol{
                          value: symbol,
                          row: index,
                          column: i
                        }
                        | symbols
                      ],
                      nil
                    }
                end
            end
        end
      end
    )
    |> then(fn {numbers, symbols, number_in_progress} ->
      case number_in_progress do
        nil ->
          {numbers, symbols}

        {value, col} ->
          {
            [
              %Number{
                value: value,
                row: index,
                column: col
              }
              | numbers
            ],
            symbols
          }
      end
    end)
  end

  def part_one do
    read_input("input.txt")
    |> Enum.with_index()
    |> Enum.map(&process_line/1)
    |> Enum.reduce(
      {[], []},
      fn {numbers, symbols}, {all_numbers, all_symbols} ->
        {all_numbers ++ numbers, all_symbols ++ symbols}
      end
    )
    |> then(fn {numbers, symbols} ->
      numbers
      |> Enum.filter(fn number ->
        symbols
        |> Enum.any?(fn symbol ->
          Symbol.is_adjacent?(symbol, number)
        end)
      end)
    end)
    |> Enum.map(& &1.value)
    |> Enum.map(&String.to_integer/1)
    |> Enum.sum()
  end

  def part_two do
    read_input("input.txt")
    |> Enum.with_index()
    |> Enum.map(&process_line/1)
    |> Enum.reduce(
      {[], []},
      fn {numbers, symbols}, {all_numbers, all_symbols} ->
        {all_numbers ++ numbers, all_symbols ++ symbols}
      end
    )
    |> then(fn {numbers, symbols} ->
      symbols
      |> Enum.filter(fn symbol ->
        symbol.value == "*"
      end)
      |> Enum.map(fn symbol ->
        numbers
        |> Enum.filter(fn number ->
          Symbol.is_adjacent?(symbol, number)
        end)
        |> then(fn adjacent_numbers ->
          {
            adjacent_numbers |> Enum.count(),
            adjacent_numbers
            |> Enum.map(& &1.value)
            |> Enum.map(&String.to_integer/1)
            |> Enum.product()
          }
        end)
      end)
      |> Enum.filter(fn {count, _} -> count == 2 end)
      |> Enum.map(&(&1 |> elem(1)))
      |> Enum.sum()
    end)
  end
end

Day03.part_one() |> IO.puts()
Day03.part_two() |> IO.puts()
