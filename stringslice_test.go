package stringslice

import (
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	type args struct {
		list []string
		item string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				list: []string{},
				item: "t3",
			},
			want: false,
		},
		{
			name: "contains_single",
			args: args{
				list: []string{"t3"},
				item: "t3",
			},
			want: true,
		},
		{
			name: "contains_multiple",
			args: args{
				list: []string{"t1", "t2", "t3", "t4"},
				item: "t3",
			},
			want: true,
		},
		{
			name: "not_contains_single",
			args: args{
				list: []string{"t1"},
				item: "t3",
			},
			want: false,
		},
		{
			name: "not_contains_multiple",
			args: args{
				list: []string{"t1", "t2", "t4"},
				item: "t3",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.list, tt.args.item); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	type args struct {
		list []string
		item string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "contains_single",
			args: args{
				list: []string{"t3"},
				item: "t3",
			},
			want: []string{},
		},
		{
			name: "contains_multiple",
			args: args{
				list: []string{"t1", "t2", "t3", "t4"},
				item: "t3",
			},
			want: []string{"t1", "t2", "t4"},
		},
		{
			name: "not_contains_empty",
			args: args{
				list: []string{},
				item: "t3",
			},
			want: []string{},
		},
		{
			name: "not_contains_multiple",
			args: args{
				list: []string{"t1", "t2", "t4"},
				item: "t3",
			},
			want: []string{"t1", "t2", "t4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Remove(tt.args.list, tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEquals(t *testing.T) {
	type args struct {
		list1 []string
		list2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				list1: []string{},
				list2: []string{},
			},
			want: true,
		},
		{
			name: "single_same",
			args: args{
				list1: []string{"test1"},
				list2: []string{"test1"},
			},
			want: true,
		},
		{
			name: "single_only_list1",
			args: args{
				list1: []string{"test1"},
				list2: []string{},
			},
			want: false,
		},
		{
			name: "single_only_list2",
			args: args{
				list1: []string{},
				list2: []string{"test1"},
			},
			want: false,
		},
		{
			name: "multiple_same",
			args: args{
				list1: []string{"test1", "test2", "test3"},
				list2: []string{"test1", "test2", "test3"},
			},
			want: true,
		},
		{
			name: "multiple_only_list1",
			args: args{
				list1: []string{"test1", "test2", "test3"},
				list2: []string{"test2"},
			},
			want: false,
		},
		{
			name: "multiple_only_list2",
			args: args{
				list1: []string{"test2"},
				list2: []string{"test1", "test2", "test3"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equals(tt.args.list1, tt.args.list2); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
