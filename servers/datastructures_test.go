package servers

import (
	"reflect"
	"strings"
	"testing"
)

func TestPush(t *testing.T) {
	l := NodeList{}
	l.Push("foo")
	l.Push("bar")
	l.Push("goo")
	if l.head.Val != "foo" {
		t.Errorf("invalid head, expected %s; got %s", "foo", l.head.Val)
	}
	if l.tail.Val != "goo" {
		t.Errorf("invalid tail, expected %s; got %s", "goo", l.tail.Val)
	}
}

func TestPop(t *testing.T) {
	l := NodeList{}
	l.Push("foo")
	n1 := l.Pop()
	if n1.Val != "foo" {
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


func TestRemove(t *testing.T) {
	l := NodeList{}
	l.Push("foo")
	l.Push("bar")
	l.Push("goo")
	n1 := l.Remove(l.Top().Next)
	if n1 != "bar" {
		t.Error("fail to remove foo")
	}
	if !reflect.DeepEqual(l.ValuesList(), []string{"foo", "goo"}) {
		t.Errorf("Incorrect lists, expected bar foo, goo, got %s", strings.Join(l.ValuesList(), ", "))
	}
	n2 := l.Remove(l.Top())
	if n2 != "foo" {
		t.Error("fail to remove foo")
	}
	if !reflect.DeepEqual(l.ValuesList(), []string{"goo"}) {
		t.Errorf("Incorrect lists, expected goo got %s", strings.Join(l.ValuesList(), ", "))
	}
}
