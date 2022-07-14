package main.java;

import java.util.Objects;

public class MyLinkedList {
	MyLinkedList next;
	int item;

	MyLinkedList(MyLinkedList next, int item) {
		this.next = next;
		this.item = item;
	}

	@Override
	public boolean equals(Object o) {
		if (this == o) {
			return true;
		}
		if (o == null || getClass() != o.getClass()) {
			return false;
		}
		MyLinkedList that = (MyLinkedList) o;
		return item == that.item && Objects.equals(next, that.next);
	}
}
