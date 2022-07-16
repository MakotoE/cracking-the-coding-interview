package main.java;

import java.util.ArrayList;
import java.util.Optional;

public class Q3_2 {
	static class StackMin {
		ArrayList<Integer> list;
		ArrayList<Integer> minimums;

		StackMin() {
			this.list = new ArrayList<>();
			this.minimums = new ArrayList<>();
		}

		void push(int item) {
			list.add(item);

			var min = min();
			if (min.isEmpty() || item < min.get()) {
				minimums.add(item);
			}
		}

		Optional<Integer> pop() {
			if (list.isEmpty()) {
				return Optional.empty();
			}

			var item = list.remove(list.size() - 1);

			//noinspection OptionalGetWithoutIsPresent
			if (min().get().equals(item)) {
				minimums.remove(minimums.size() - 1);
			}

			return Optional.of(item);
		}

		Optional<Integer> min() {
			if (minimums.isEmpty()) {
				return Optional.empty();
			}

			return Optional.of(minimums.get(minimums.size() - 1));
		}
	}

	public static void main(String[] args) {
		var stack = new StackMin();
		assert(stack.pop().isEmpty());
		assert(stack.min().isEmpty());

		stack.push(0);
		assert(stack.min().get() == 0);
		assert(stack.pop().get() == 0);
		assert(stack.pop().isEmpty());

		stack.push(1);
		stack.push(0);
		assert(stack.min().get() == 0);
		assert(stack.pop().get() == 0);
		assert(stack.min().get() == 1);
		assert(stack.pop().get() == 1);
		assert(stack.pop().isEmpty());

		stack.push(0);
		stack.push(1);
		assert(stack.min().get() == 0);
		assert(stack.pop().get() == 1);
		assert(stack.min().get() == 0);
		assert(stack.pop().get() == 0);
		assert(stack.pop().isEmpty());
	}
}
