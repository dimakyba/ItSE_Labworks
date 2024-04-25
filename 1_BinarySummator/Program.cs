using System;
using System.Text;

namespace BinarySummator
{
  class Program
  {
    static void Main(string[] args)
    {
      Console.OutputEncoding = UTF8Encoding.UTF8;

      Console.WriteLine("Introduction to Software Engineering, Labwork #1\nVariant 7 mod 5 = 2\nLength of combination - 7 symbols.\n");

      bool isValidInput = false;
      while (!isValidInput)
      {
        Console.Write("Enter the first binary number : ");
        string firstBinary = Console.ReadLine();
        Console.Write("Enter the second binary number: ");
        string secondBinary = Console.ReadLine();
        Console.WriteLine();

        if (ValidateBinary(firstBinary) && ValidateBinary(secondBinary))
        {
          Console.WriteLine($"The first binary number in decimal representation: {ConvertBinaryToDecimal(firstBinary)}");
          Console.WriteLine($"The second binary number in decimal representation: {ConvertBinaryToDecimal(secondBinary)}");
          string sum = GetSumOfTwoBinaries(firstBinary, secondBinary);
          Console.WriteLine($"The sum of two binary numbers is {sum}\nThe decimal representation of the sum is {ConvertBinaryToDecimal(sum)}");
          isValidInput = true;
        }
        else
        {
          Console.WriteLine("The provided input is incorrect, please try again.");
        }
      }
    }


    static string GetSumOfTwoBinaries(string bin1, string bin2)
    {
      string result = "";
      int sum = 0;
      int i = bin1.Length - 1, j = bin2.Length - 1;
      //створюється порожній рядок result і дві змінні sum і i, j, що вказують на останній символ кожного рядка.

      while (i >= 0 || j >= 0 || sum == 1)
      // Далі використовується цикл while, який буде продовжуватися, поки є цифри у двох введених рядках або доданок sum дорівнює 1.
      {
        if (j >= 0)
        // Перевіряється, чи є ще цифри у bin1 або bin2 (i >= 0 або j >= 0).
        // Якщо j >= 0, то до змінної sum додається значення поточного біта з bin1 та bin2, але вони перетворюються у цілі числа, віднімаючи значення '0' (ASCII-код нуля) двічі. Це робиться для того, щоб отримати фактичне значення числа з рядка. Наприклад, '1' - '0' дасть 1, а '0' - '0' дасть 0.
        {
          sum += bin1[i] + bin2[j] - 2 * '0';

        }
        // До результуючого рядка result додається символ, який відповідає залишковому від числа sum при діленні на 2 (це дасть наступний біт у відповіді).
        // sum ділиться на 2 для того, щоб підготуватися до наступної ітерації. Індекси i та j зменшуються, щоб перейти до наступних бітів у bin1 та bin2.
        result += (char)(sum % 2 + '0');
        sum /= 2;
        i--; j--;
      }

      return result;
    }

    static int ConvertBinaryToDecimal(string binary)
    {
      int result = 0;
      for (int i = 0; i < binary.Length; i++)
      {
        result += ((int)binary[i] - 48) * (int)Math.Pow(2, binary.Length - 1 - i);
      }

      return result;
    }


    static bool ValidateBinary(string binary)
    {
      for (int i = 0; i < binary.Length; i++)
      {
        if (binary.Length != 7) //hardcode
        {
          Console.WriteLine("That length is not 7");
          return false;
        }
        if (binary[i] != '0' && binary[i] != '1')
        {
          return false;
        }
      }

      return true;
    }
  }
}
