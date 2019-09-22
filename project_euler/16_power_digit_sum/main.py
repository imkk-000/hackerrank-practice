from functools import reduce
if __name__ == '__main__':
    for i in range(int(input())):
        print(reduce(
            lambda acc, n: acc + int(n),
            str(1 << int(input())), 0))
