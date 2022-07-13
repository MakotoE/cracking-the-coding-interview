namespace c_sharp
{
    [TestClass]
    public class Question1_2
    {
        static bool IsPermutation(string a, string b)
        {
            var map = new int[0xa0];

            foreach (var c in a)
            {
                map[c] += 1;
            }

            foreach (var c in b)
            {
                map[c] -= 1;
            }

            foreach (var frequency in map)
            {
                if (frequency != 0)
                {
                    return false;
                }
            }

            return true;
        }

        [TestMethod]
        public void TestMethod1()
        {
            var tests = new (string a, string b, bool expected)[]
            {
                (
                    "",
                    "",
                    true
                ),
                (
                    "a",
                    "",
                    false
                ),
                (
                    "",
                    "a",
                    false
                ),
                (
                    "a",
                    "a",
                    true
                ),
                (
                    "aa",
                    "a",
                    false
                ),
                (
                    "a",
                    "aa",
                    false
                ),
                (
                    "ab",
                    "ba",
                    true
                ),
                (
                    "aba",
                    "baa",
                    true
                ),
                (
                    "aba",
                    "baaa",
                    false
                ),
            };

            for (var i = 0; i < tests.Length; i++)
            {
                var result = IsPermutation(tests[i].a, tests[i].b);
                Assert.AreEqual(tests[i].expected, result, i.ToString());
            }
        }
    }
}