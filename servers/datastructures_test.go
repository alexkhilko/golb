package servers

import (
	"testing"
	"reflect"
	"strings"
)

func TestPush(t *testing.T) {
	l := NodeList{}
	l.Push("foo")
	l.Push("bar")
	l.Push("goo")
	if l.head.val != "foo" {
		t.Errorf("invalid head, expected %s; got %s", "foo", l.head.val)
	}
	if l.tail.val != "goo" {
		t.Errorf("invalid tail, expected %s; got %s", "goo", l.tail.val)
	}
}

func TestPop(t *testing.T) {
	l := NodeList{}
	l.Push("foo")
	n1 := l.Pop()
	if n1.val != "foo"  {
		t.Errorf("fail on pop, expected foo; got %s", n1)
	}
	n2 := l.Pop()
	if n2 != nil {
		t.Errorf("fail on empty pop, expected nil; got %s", n2)
	}
	if l.head != nil || l.tail != nil {
		t.Errorf("expected empty list; got %s head, %s tail", l.head, l.tail)
	}
}


func toValueList(nl NodeList) []string {
	if nl.head == nil {
		return []string{}
	}
	cur := nl.head
	l := []string{}
	for cur != nil {
		l = append(l, cur.val)
		cur = cur.next
	}
	return l
}

func TestRemove(t *testing.T) {
	l := NodeList{}
	l.Push("foo")
	l.Push("bar")
	n1 := l.Remove("foo")
	if n1 != "foo"  {
		t.Error("fail to remove foo")
	}
	if reflect.DeepEqual(toValueList(l), []string{"bar", "goo"}) {
		t.Errorf("Incorrect lists, expected bar got %s", strings.Join(toValueList(l), ", "))
	}
	n2 := l.Remove("bar")
	if n2 != "bar"  {
		t.Error("fail to remove bar")
	}
	if reflect.DeepEqual(toValueList(l), []string{}) {
		t.Errorf("Incorrect lists, expected bar, got %s", strings.Join(toValueList(l), ", "))
	}
	n3 := l.Remove("empty")
	if n3 != ""  {
		t.Error("fail to remove empty")
	}

}