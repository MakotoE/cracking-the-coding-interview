package main.java;

import java.util.Optional;

public class Q2_2 {
	// Returns nth to last element in list if it exists.
	static Optional<Integer> GetNthToLast(MyLinkedList list, int n) {
		var curr = list;
		int size = 0;
		while (curr != null) {
			curr = curr.next;
			size++;
		}

		if (n >= size) {
			return Optional.empty();
		}

		curr = list;
		for (int i = 0; i < size - 1 - n; i++) {
			curr = curr.next;
		}

		return Optional.of(curr.item);
	}

	static class TestCase {
		MyLinkedList list;
		int n;
		Optional<Integer> expected;

		TestCase(MyLinkedList list, int n, Optional<Integer> expected) {
			this.list = list;
			this.n = n;
			this.expected = expected;
		}
	}

	public static void main(String[] args) {
		var tests = new TestCase[] {
			new TestCase(
				new MyLinkedList(null, 0),
				0,
				Optional.of(0)
			),
			new TestCase(
				new MyLinkedList(null, 0),
				1,
				Optional.empty()
			),
			new TestCase(
				new MyLinkedList(new MyLinkedList(null, 1), 0),
				0,
				Optional.of(1)
			),
			new TestCase(
				new MyLinkedList(new MyLinkedList(null, 1), 0),
				1,
				Optional.of(0)
			),
			new TestCase(
				new MyLinkedList(new MyLinkedList(null, 1), 0),
				2,
				Optional.empty()
			),
		};

		for (int i = 0; i < tests.length; i++) {
			var result = GetNthToLast(tests[i].list, tests[i].n);
			assert(result.equals(tests[i].expected)) : i;
		}
	}
}
