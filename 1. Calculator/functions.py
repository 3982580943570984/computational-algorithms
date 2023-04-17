def sqrt(a: float, eps: float):
    if (a <= 0):
        return 0

    x = a
    x0 = 0.0
    while (abs(x - x0) > eps):
        x0 = x
        x = (x0 + a / x0) / 2
    return x


def sin(x: float, eps: float):
    ''' x == radians '''
    result = x
    k = 1
    u0 = x
    u1 = -((x ** 2) / (2 * k * (2 * k + 1))) * u0
    while (abs(u1) > eps):
        result += u1
        k += 1
        u0 = u1
        u1 = -((x ** 2) / (2 * k * (2 * k + 1))) * u0
    return result


def cos(x: float, eps: float):
    ''' x == radians '''
    result = 1
    k = 1
    u0 = 1
    u1 = -((x ** 2) / (2 * k * (2 * k - 1))) * u0
    while (abs(u1) > eps):
        result += u1
        k += 1
        u0 = u1
        u1 = -((x ** 2) / (2 * k * (2 * k - 1))) * u0
    return result


def tan(x: float, eps: float):
    return (sin(x=x, eps=eps) / cos(x=x, eps=eps))


def ctg(x: float, eps: float):
    # return (1 / tan(x, eps))
    return (cos(x=x, eps=eps) / sin(x=x, eps=eps))


def e(x: float, eps: float):
    result = 1
    k = 1
    u0 = 1
    u1 = (x / k) * u0
    while (abs(u1) > eps):
        result += u1
        k += 1
        u0 = u1
        u1 = (x / k) * u0
    return result


def ln(x: float, eps: float):
    if (x <= 0):
        return 0

    result = (1 - x) / (1 + x)
    k = 1
    u0 = (1 - x) / (1 + x)
    u1 = (u0 ** (2*k + 1)) / (2 * k + 1)
    while (abs(u1) > eps):
        result += u1
        k += 1
        u1 = (u0 ** (2*k + 1)) / (2 * k + 1)
    return result * -2


if __name__ == "__main__":
    print(e(4, 0.0000001))

    print(sin(3, 0.0000001))
    print(cos(3, 0.0000001))
    print(tan(3, 0.0000001))
    print(ctg(3, 0.0000001))

    print(sqrt(3, 0.0000001))

    print(ln(4, 0.000000000000001))
