package main.java;

import java.util.Optional;

public class Q2_4 {
	// Returns list partitioned by given value so that all values greater than equal to value come
	// after other items.
	static ListNode Partition(ListNode list, int value) {
		Optional<ListNode> head = Optional.empty();
		Optional<ListNode> headEnd = Optional.empty();
		Optional<ListNode> partition = Optional.empty();
		Optional<ListNode> partitionEnd = Optional.empty();

		var curr = list;
		while (curr != null) {
			if (curr.item < value) {
				if (headEnd.isEmpty()) {
					head = Optional.of(new ListNode(null, curr.item));
					headEnd = head;
				} else {
					headEnd.get().next = new ListNode(null, curr.item);
					headEnd = Optional.of(headEnd.get().next);
				}
			} else {
				if (partitionEnd.isEmpty()) {
					partition = Optional.of(new ListNode(null, curr.item));
					partitionEnd = partition;
				} else {
					partitionEnd.get().next = new ListNode(null, curr.item);
					partitionEnd = Optional.of(partitionEnd.get().next);
				}
			}

			curr = curr.next;
		}

		if (partition.isPresent()) {
			if (head.isPresent()) {
				headEnd.get().next = partition.get();
			} else {
				head = partition;
			}
		}

		return head.orElseThrow();
	}

	static class TestCase {
		ListNode list;
		int value;
		ListNode expected;

		TestCase(ListNode list, int value, ListNode expected) {
			this.list = list;
			this.value = value;
			this.expected = expected;
		}
	}

	public static void main(String[] args) {
		var tests = new TestCase[] {
			new TestCase(
				new ListNode(null, 0),
				0,
				new ListNode(null, 0)
			),
			new TestCase(
				new ListNode(new ListNode(null, 1), 0),
				0,
				new ListNode(new ListNode(null, 1), 0)
			),
			new TestCase(
				new ListNode(new ListNode(null, 1), 0),
				1,
				new ListNode(new ListNode(null, 1), 0)
			),
			new TestCase(
				new ListNode(new ListNode(null, 0), 1),
				0,
				new ListNode(new ListNode(null, 0), 1)
			),
			new TestCase(
				new ListNode(new ListNode(null, 0), 1),
				1,
				new ListNode(new ListNode(null, 1), 0)
			),
			new TestCase(
				new ListNode(new ListNode(new ListNode(null, 0), 1), 2),
				1,
				new ListNode(new ListNode(new ListNode(null, 1), 2), 0)
			),
		};

		for (int i = 0; i < tests.length; i++) {
			var result = Partition(tests[i].list, tests[i].value);
			assert(result.equals(tests[i].expected)) : i;
		}
	}
}
