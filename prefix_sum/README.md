# What is Prefix Sum Array?
  A prefix sum array, also known as a cumulative sum array, is a derived
array that stores the cumulative sum of elements in a given array. Each
element in the prefix sum array represents the sum of all elements up to
that index in the original array. It acts as a precursor to answering
queries related to cumulative sums, allowing for fast and efficient
computations. It also reduces time complexity giving us a way out of TLE.

## Constructing Prefix Sum Array
  To construct a one-dimensional prefix sum array, we iterate through the
input array and compute the cumulative sum up to each index. The prefix
sum value at index i is the sum of all elements from index 0 to i in the
original array. Here are the steps involved in the construction of the
prefix sum array:

1. Initialize an array, say prefixSum[], of the same length as the input array.
2. Set the first element of the prefixSum[] as the same as the first element of the input array.
3. Iterate through the input array starting from the second element.
4. For each element, calculate the prefix sum by adding the current element with the prefix sum value of the previous index.
5. Assign the prefix sum value to the corresponding index in the prefixSum array.
6. Repeat steps 3–5 until the end of the input array is reached.
