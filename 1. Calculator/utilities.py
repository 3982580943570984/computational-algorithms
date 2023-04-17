from dearpygui import dearpygui as dpg
from constants import *
from functions import *


'''First row'''


def make_eps(sender, data, user_data):
    compute_widget = user_data[0]

    entry_widget = user_data[1]
    entry = float(dpg.get_value(entry_widget))

    precision_widget = user_data[2]
    precision = float(dpg.get_value(precision_widget))

    dpg.set_value(compute_widget, e(entry, precision))


def clear_entry(sender, data, user_data):
    widget = user_data
    dpg.set_value(widget, 0)


def clear_all(sender, data, user_data):
    entry_widget = user_data[0]
    compute_widget = user_data[1]
    dpg.set_value(entry_widget, 0)
    dpg.set_value(compute_widget, '0')


def delete_last_digit(sender, data, user_data):
    entry_widget = user_data[1]
    entry = float(dpg.get_value(entry_widget))

    if entry >= 0:
        dpg.set_value(entry_widget, entry // 10)
    else:
        dpg.set_value(entry_widget, entry // 10 + 1)


''' Second row '''


def make_natural_logarithm(sender, data, user_data):
    compute_widget = user_data[0]

    entry_widget = user_data[1]
    entry = float(dpg.get_value(entry_widget))

    precision_widget = user_data[2]
    precision = float(dpg.get_value(precision_widget))

    dpg.set_value(compute_widget, ln(entry, precision))


def make_inverse(sender, data, user_data):
    widget = user_data[1]
    number = float(dpg.get_value(user_data[0]))
    if number != 0:
        dpg.set_value(widget, float(1.0 / number))


def make_pow(sender, data, user_data):
    widget = user_data[1]
    number = float(dpg.get_value(user_data[0]))
    dpg.set_value(widget, float(pow(number, 2)))


def make_sqrt(sender, data, user_data):
    compute_widget = user_data[0]

    entry_widget = user_data[1]
    entry = float(dpg.get_value(entry_widget))

    precision_widget = user_data[2]
    precision = float(dpg.get_value(precision_widget))

    dpg.set_value(compute_widget, sqrt(entry, precision))


'''Third row'''


def make_sin(sender, data, user_data):
    compute_widget = user_data[0]

    entry_widget = user_data[1]
    entry = float(dpg.get_value(entry_widget))

    precision_widget = user_data[2]
    precision = float(dpg.get_value(precision_widget))

    dpg.set_value(compute_widget, sin(entry, precision))


'''Fourth row'''


def make_cos(sender, data, user_data):
    compute_widget = user_data[0]

    entry_widget = user_data[1]
    entry = float(dpg.get_value(entry_widget))

    precision_widget = user_data[2]
    precision = float(dpg.get_value(precision_widget))

    dpg.set_value(compute_widget, cos(entry, precision))


'''Fifth row'''


def make_tan(sender, data, user_data):
    compute_widget = user_data[0]

    entry_widget = user_data[1]
    entry = float(dpg.get_value(entry_widget))

    precision_widget = user_data[2]
    precision = float(dpg.get_value(precision_widget))

    dpg.set_value(compute_widget, tan(entry, precision))


'''Sixth row'''


def make_ctg(sender, data, user_data):
    compute_widget = user_data[0]

    entry_widget = user_data[1]
    entry = float(dpg.get_value(entry_widget))

    precision_widget = user_data[2]
    precision = float(dpg.get_value(precision_widget))

    dpg.set_value(compute_widget, ctg(entry, precision))


def make_negative(sender, data, user_data):
    entry_widget = user_data[1]
    entry = float(dpg.get_value(entry_widget))

    dpg.set_value(entry_widget, -entry)


def make_float(sender, data, user_data):
    entry_widget = user_data[1]
    entry = float(dpg.get_value(entry_widget))
    dpg.set_value(entry_widget, entry)


'''Miscellaneous'''


def add_digit(sender, data, user_data):
    widget = user_data[0]
    number = float(dpg.get_value(widget))
    if number == 0:
        dpg.set_value(widget, user_data[1])
    elif number > 0:
        dpg.set_value(widget, number * 10 + user_data[1])
    else:
        dpg.set_value(widget, number * 10 - user_data[1])
