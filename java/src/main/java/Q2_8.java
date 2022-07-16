package main.java;

import java.util.Optional;

public class Q2_8 {
	// Returns the item of the node where a loop begins.
	static Optional<Integer> DetectLoop(ListNode list) {
		var slow = list;
		var fast = list;

		while (slow != null && fast != null && fast.next != null) {
			slow = slow.next;
			fast = fast.next.next;

			if (slow == fast) {
				return Optional.of(slow.item);
			}
		}

		return Optional.empty();
	}

	static class TestCase {
		ListNode list;
		Optional<Integer> expected;

		TestCase(ListNode list, Optional<Integer> expected) {
			this.list = list;
			this.expected = expected;
		}
	}

	public static void main(String[] args) {
		var loopA = new ListNode(null, 0);
		loopA.next = loopA;

		var loopB = new ListNode(new ListNode(null, 1), 0);
		loopB.next.next = loopB;

		var loopC = new ListNode(new ListNode(null, 1), 0);
		loopC.next.next = loopC.next;

		var loopD = new ListNode(new ListNode(new ListNode(null, 2), 1), 0);
		loopD.next.next.next = loopD.next;

		var loopE = new ListNode(new ListNode(new ListNode(null, 2), 1), 0);
		loopE.next.next.next = loopE.next.next;

		var tests = new TestCase[] {
			new TestCase(
				new ListNode(null, 0),
				Optional.empty()
			),
			new TestCase(
				new ListNode(new ListNode(null, 1), 0),
				Optional.empty()
			),
			new TestCase(
				new ListNode(new ListNode(new ListNode(null, 2), 1), 0),
				Optional.empty()
			),
			new TestCase(
				loopA,
				Optional.of(0)
			),
			new TestCase(
				loopB,
				Optional.of(0)
			),
			new TestCase(
				loopC,
				Optional.of(1)
			),
			new TestCase(
				loopD,
				Optional.of(2)
			),
			new TestCase(
				loopE,
				Optional.of(2)
			),
		};

		for (int i = 0; i < tests.length; i++) {
			var result = DetectLoop(tests[i].list);
			assert(result.equals(tests[i].expected)) : i;
		}
	}
}
