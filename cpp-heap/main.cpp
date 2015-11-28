#include <array>
#include <cstdio>
#include <cassert>

std::array<int, 10> g_data;
int g_size = 0;

void swap(int *a, int *b)
{
	int tmp = *a;
	*a = *b;
	*b = tmp;
}

void print_heap()
{
	for(int i = 0 ; i < g_size ; ++i) {
		printf("%d ", g_data[i]);
	}
	printf("\n");
}

void push_heap(int val)
{
	int curr_idx = g_size;
	g_data[curr_idx] = val;

	while(curr_idx != 0) {
		int parent_idx = (curr_idx - 1) / 2;
		if(g_data[parent_idx] > g_data[curr_idx]) {
			swap(&g_data[parent_idx], &g_data[curr_idx]);
			curr_idx = parent_idx;
		} else {
			break;
		}
	}

	g_size += 1;
}

int pop_heap()
{
	int retval = g_data[0];
	g_data[0] = g_data[g_size - 1];
	g_size -= 1;

	// fix heap
	int curr_idx = 0;
	while(curr_idx < g_size) {
		int left_idx = (curr_idx * 2) + 1;
		int right_idx = (curr_idx * 2) + 2;

		int curr = g_data[curr_idx];
		int left = g_data[left_idx];
		int right = g_data[right_idx];

		if(left_idx >= g_size) {
			left = 999;
		}
		if(right_idx >= g_size) {
			right = 999;
		}

		if(curr <= left && curr <= right) {
			break;
		}

		if(left <= curr && left <= right) {
			swap(&g_data[curr_idx], &g_data[left_idx]);
			curr_idx = left_idx;
		}
		else if(right <= curr && right <= left) {
			swap(&g_data[curr_idx], &g_data[right_idx]);
			curr_idx = right_idx;
		}
		else
		{
			assert(!"do not reach");
		}
	}
	return retval;
}


int main()
{
	push_heap(6);
	print_heap();
	push_heap(4);
	print_heap();
	push_heap(9);
	print_heap();
	push_heap(7);
	print_heap();
	push_heap(1);
	print_heap();
	push_heap(2);
	print_heap();

	assert(1 == pop_heap());
	print_heap();
	assert(2 == pop_heap());
	print_heap();
	assert(4 == pop_heap());
	print_heap();
	assert(6 == pop_heap());
	print_heap();
	assert(7 == pop_heap());
	print_heap();
	assert(9 == pop_heap());
	print_heap();

	return 0;
}
