#include <cstdio>

int process(int argc, char *argv[]) {
  fprintf(stdout, "Args (%d):\n", argc);
  for (int i = 0; i < argc; ++i) {
    fprintf(stdout, "%s\n", argv[i]);
  }
  return 0;
}
