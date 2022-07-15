package main.java;

import java.util.HashSet;
import java.util.List;

public class Q2_1 {
	// Removes duplicates from list.
	static void RemoveDuplicates(ListNode list) {
		HashSet<Integer> set = new HashSet<>(List.of(list.item));

		var curr = list;
		while (curr.next != null) {
			if (set.contains(curr.next.item)) {
				curr.next = curr.next.next;
			} else {
				set.add(curr.next.item);
				curr = curr.next;
			}
		}
	}

	static class TestCase {
		ListNode list;
		ListNode expected;

		TestCase(ListNode list, ListNode expected) {
			this.list = list;
			this.expected = expected;
		}
	}

	public static void main(String[] args) {
		var tests = new TestCase[] {
			new TestCase(
				new ListNode(null, 0),
				new ListNode(null, 0)
			),
			new TestCase(
				new ListNode(new ListNode(null, 1), 0),
				new ListNode(new ListNode(null, 1), 0)
			),
			new TestCase(
				new ListNode(new ListNode(null, 0), 0),
				new ListNode(null, 0)
			),
			new TestCase(
				new ListNode(
					new ListNode(new ListNode(null, 0), 1),
					0
				),
				new ListNode(new ListNode(null, 1), 0)
			),
			new TestCase(
				new ListNode(
					new ListNode(
						new ListNode(new ListNode(null, 1), 0),
						1
					),
					0
				),
				new ListNode(new ListNode(null, 1), 0)
			),
		};

		for (int i = 0; i < tests.length; i++) {
			RemoveDuplicates(tests[i].list);
			assert(tests[i].list.equals(tests[i].expected)) : i;
		}
	}
}
