package main.java;

public class Question1_9 {
	// Returns true if a is a rotation of b.
	static boolean IsRotation(String a, String b) {
		if (a.length() != b.length()) {
			return false;
		}

		String rotated = a;
		for (int i = 0; i < a.length(); i++) {
			if (rotated.equals(b)) {
				return true;
			}

			rotated = rotated.substring(1) + rotated.charAt(0);
		}

		return a.length() == 0;
	}

	public static void main(String[] args) {
		assert(IsRotation("", ""));
		assert(!IsRotation("a", ""));
		assert(!IsRotation("", "a"));
		assert(IsRotation("a", "a"));
		assert(!IsRotation("aa", "a"));
		assert(IsRotation("aa", "aa"));
		assert(IsRotation("ab", "ab"));
		assert(IsRotation("ab", "ba"));
		assert(IsRotation("abc", "bca"));
		assert(!IsRotation("abc", "bac"));
		assert(!IsRotation("abc", "cab"));
	}
}
