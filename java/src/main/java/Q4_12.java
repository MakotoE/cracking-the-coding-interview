package main.java;

public class Q4_12 {
	// Returns number of paths starting from root where its items sum to given value.
	static int CountPathsWithSum(BinaryNode root, int value) {
		int result = 0;

		if (root.item == value) {
			result = 1;
		}

		if (root.left != null) {
			result += CountPathsWithSum(root.left, value - root.item);
		}
		if (root.right != null) {
			result += CountPathsWithSum(root.right, value - root.item);
		}

		return result;
	}

	static class TestCase {
		BinaryNode root;
		int value;
		int expected;

		TestCase(BinaryNode root, int value, int expected) {
			this.root = root;
			this.value = value;
			this.expected = expected;
		}
	}

	public static void main(String[] args) {
		var tests = new TestCase[] {
			new TestCase(
				new BinaryNode(0, null, null),
				0,
				1
			),
			new TestCase(
				new BinaryNode(0, null, null),
				1,
				0
			),
			new TestCase(
				new BinaryNode(1, null, null),
				0,
				0
			),
			new TestCase(
				new BinaryNode(
					0,
					new BinaryNode(0, null, null),
					null
				),
				0,
				2
			),
			new TestCase(
				new BinaryNode(
					0,
					new BinaryNode(1, null, null),
					null
				),
				0,
				1
			),
			new TestCase(
				new BinaryNode(
					0,
					new BinaryNode(1, null, null),
					null
				),
				1,
				1
			),
			new TestCase(
				new BinaryNode(
					0,
					new BinaryNode(0, null, null),
					new BinaryNode(0, null, null)
				),
				0,
				3
			),
			new TestCase(
				new BinaryNode(
					0,
					new BinaryNode(1, null, null),
					new BinaryNode(1, null, null)
				),
				1,
				2
			),
			new TestCase(
				new BinaryNode(
					1,
					new BinaryNode(1, null, null),
					new BinaryNode(1, null, null)
				),
				2,
				2
			),
			new TestCase(
				new BinaryNode(
					1,
					new BinaryNode(3, null, null),
					new BinaryNode(1, null, null)
				),
				2,
				1
			),
			new TestCase(
				new BinaryNode(
					1,
					new BinaryNode(
						1,
						new BinaryNode(1, null, null),
						null
					),
					new BinaryNode(1, null, null)
				),
				3,
				1
			),
			new TestCase(
				new BinaryNode(
					1,
					new BinaryNode(
						3,
						new BinaryNode(9, null, null),
						null
					),
					new BinaryNode(
						2,
						new BinaryNode(5, null, null),
						new BinaryNode(
							7,
							new BinaryNode(3, null, null),
							null
						)
					)
				),
				13,
				2
			),
		};

		for (int i = 0; i < tests.length; i++) {
			var result = CountPathsWithSum(tests[i].root, tests[i].value);
			assert(result == tests[i].expected) : i;
		}
	}
}
