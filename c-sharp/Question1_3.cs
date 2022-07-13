namespace c_sharp
{
    [TestClass]
    public class Question1_3
    {
        // Encodes spaces in s to "%20".
        void URLify(char[] s)
        {
            for (int i = 0; i < s.Length; i++)
            {
                if (s[i] == ' ' && i + 2 < s.Length)
                {
                    s[i] = '%';
                    s[i + 1] = '2';
                    s[i + 2] = '0';
                }
            }
        }

        [TestMethod]
        public void TestURLify()
        {
            var tests = new (string s, string expected)[]
            {
                (
                    "",
                    ""
                ),
                (
                    "a",
                    "a"
                ),
                (
                    " ",
                    " "
                ),
                (
                    "   ",
                    "%20"
                ),
                (
                    "a   ",
                    "a%20"
                ),
                (
                    "   a",
                    "%20a"
                ),
                (
                    "a   a",
                    "a%20a"
                ),
                (
                    "      ",
                    "%20%20"
                ),
                (
                    "   a   ",
                    "%20a%20"
                ),
                (
                    "    ",
                    "%20 "
                ),
            };

            for (int i = 0; i < tests.Length; i++)
            {
                var s = tests[i].s.ToCharArray();
                URLify(s);
                CollectionAssert.AreEqual(tests[i].expected.ToCharArray(), s);
            }
        }
    }
}
