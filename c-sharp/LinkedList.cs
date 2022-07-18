namespace c_sharp
{
    class LinkedList
    {
        class Node
        {
            public int item;
            public Node? next;

            public Node(int item, Node? next)
            {
                this.item = item;
                this.next = next;
            }
        }

        Node? head;
        int size;

        public LinkedList()
        {
            head = null;
            size = 0;
        }

#pragma warning disable CS8602 // Dereference of a possibly null reference.
        public void Add(int index, int item)
        {
            if (index == size + 1)
            {
                throw new IndexOutOfRangeException();
            }

            if (index == 0)
            {
                head = new Node(item, head);
            }
            else
            {

                var parent = head;
                for (int i = 0; i < index - 1; i++)
                {
                    parent = parent.next;
                }

                parent.next = new Node(item, parent.next);
            }

            size++;
        }

        public int Get(int index)
        {
            if (index >= size)
            {
                throw new IndexOutOfRangeException();
            }

            var curr = head;
            for (int i = 0; i < index; i++)
            {
                curr = curr.next;
            }
            return curr.item;
        }

        public int Size()
        {
            return size;
        }

        public void Remove(int index)
        {
            if (index >= size)
            {
                throw new IndexOutOfRangeException();
            }

            if (index == 0)
            {
                head = head.next;
            }
            else
            {
                var parent = head;
                for (int i = 0; i < index - 1; i++)
                {
                    parent = parent.next;
                }
                parent.next = parent.next.next;
            }
            size--;
        }
#pragma warning restore CS8602 // Dereference of a possibly null reference.
    }

    [TestClass]
    public class LinkedListTest
    {
        [TestMethod]
        public void TestAdd()
        {
            var list = new LinkedList();
            Assert.ThrowsException<IndexOutOfRangeException>(() =>
            {
                list.Add(1, 0);
            });
            Assert.ThrowsException<IndexOutOfRangeException>(() =>
            {
                list.Get(0);
            });
            Assert.AreEqual(0, list.Size());

            list.Add(0, 0);
            Assert.AreEqual(0, list.Get(0));

            list.Add(1, 1);
            Assert.AreEqual(0, list.Get(0));
            Assert.AreEqual(1, list.Get(1));

            list.Add(0, 2);
            Assert.AreEqual(2, list.Get(0));
            Assert.AreEqual(0, list.Get(1));
            Assert.AreEqual(1, list.Get(2));

            list.Add(1, 3);
            Assert.AreEqual(2, list.Get(0));
            Assert.AreEqual(3, list.Get(1));
            Assert.AreEqual(0, list.Get(2));
            Assert.AreEqual(1, list.Get(3));
            Assert.AreEqual(4, list.Size());
        }

        [TestMethod]
        public void TestRemove()
        {
            var list = new LinkedList();
            Assert.ThrowsException<IndexOutOfRangeException>(() =>
            {
                list.Remove(0);
            });

            list.Add(0, 0);
            Assert.AreEqual(1, list.Size());
            Assert.AreEqual(0, list.Get(0));

            list.Remove(0);
            Assert.AreEqual(0, list.Size());

            list.Add(0, 0);
            list.Add(1, 1);
            Assert.AreEqual(2, list.Size());
            Assert.AreEqual(0, list.Get(0));
            Assert.AreEqual(1, list.Get(1));

            list.Remove(1);
            Assert.AreEqual(1, list.Size());
            Assert.AreEqual(0, list.Get(0));

            list.Add(1, 1);
            list.Remove(0);
            Assert.AreEqual(1, list.Size());
            Assert.AreEqual(1, list.Get(0));

            list.Remove(0);
            Assert.AreEqual(0, list.Size());

            list.Add(0, 0);
            list.Add(1, 1);
            list.Add(2, 2);

            list.Remove(1);
            Assert.AreEqual(2, list.Size());
            Assert.AreEqual(0, list.Get(0));
            Assert.AreEqual(2, list.Get(1));
        }
    }
}
