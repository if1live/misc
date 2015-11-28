template
struct factorial {
  enum { value = factorial::value * N };
};
template<>
struct factorial<1> {
  enum { value = 1 };
};
