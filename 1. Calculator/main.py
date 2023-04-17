from dearpygui import dearpygui as dpg
from utilities import *
from constants import *

dpg.create_context()
dpg.setup_dearpygui()
dpg.create_viewport(title='Calculator', width=SCREEN_WIDTH,
                    height=SCREEN_HEIGHT, x_pos=DISPLAY_CENTER_X, y_pos=DISPLAY_CENTER_Y,
                    small_icon='./icons/llvm.ico', large_icon='./icons/llvm.ico')

with dpg.window(tag='Primary', width=SCREEN_WIDTH, height=SCREEN_HEIGHT,
                min_size=[300, 350], pos=[0, 0], no_close=True, no_move=True,
                no_collapse=True, no_title_bar=True) as main_window:

    with dpg.font_registry():
        with dpg.font(file="C:/Windows/Fonts/JetBrains Mono Thin Nerd Font Complete Mono.ttf", size=15, parent=dpg.last_item()) as default_font:
            dpg.add_font_range_hint(hint=dpg.mvFontRangeHint_Cyrillic)

    dpg.bind_font(default_font)

    dpg.add_text('Запись вычисления')
    compute_widget = dpg.add_input_text(
        tag='Compute', width=SCREEN_WIDTH - 32, default_value='0', readonly=True)
    dpg.add_text('Текущая запись')
    entry_widget = dpg.add_input_text(
        tag='Entry', width=SCREEN_WIDTH - 32, default_value='0', scientific=True)
    dpg.add_text('Точность вычисления')
    precision_widget = dpg.add_input_text(
        tag='Precision', width=SCREEN_WIDTH - 32, default_value='0.000001', decimal=True)

    # dpg.add_item_clicked_handler -
    # Добавить графу - точность вычисления

    with dpg.group(horizontal=True):
        dpg.add_button(label='e^x', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=make_eps, user_data=[compute_widget, entry_widget, precision_widget])
        dpg.add_button(label='CE', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=clear_entry, user_data=entry_widget)
        dpg.add_button(label='C', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=clear_all, user_data=[entry_widget, compute_widget])
        dpg.add_button(label='DEL', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=delete_last_digit, user_data=[compute_widget, entry_widget, precision_widget])

    with dpg.group(horizontal=True):
        dpg.add_button(label='ln(x)', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=make_natural_logarithm, user_data=[compute_widget, entry_widget, precision_widget])
        dpg.add_button(label='1/x', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=make_inverse, user_data=[entry_widget, compute_widget])
        dpg.add_button(label='x^2', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=make_pow, user_data=[entry_widget, compute_widget])
        dpg.add_button(label='sqrt(x)', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=make_sqrt, user_data=[compute_widget, entry_widget, precision_widget])

    with dpg.group(horizontal=True):
        dpg.add_button(label='sin(x)', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=make_sin, user_data=[compute_widget, entry_widget, precision_widget])
        dpg.add_button(label='7', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=add_digit, user_data=[entry_widget, 7])
        dpg.add_button(label='8', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=add_digit, user_data=[entry_widget, 8])
        dpg.add_button(label='9', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=add_digit, user_data=[entry_widget, 9])

    with dpg.group(horizontal=True):
        dpg.add_button(label='cos(x)', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=make_cos, user_data=[compute_widget, entry_widget, precision_widget])
        dpg.add_button(label='4', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=add_digit, user_data=[entry_widget, 4])
        dpg.add_button(label='5', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=add_digit, user_data=[entry_widget, 5])
        dpg.add_button(label='6', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=add_digit, user_data=[entry_widget, 6])

    with dpg.group(horizontal=True):
        dpg.add_button(label='tan(x)', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=make_tan, user_data=[compute_widget, entry_widget, precision_widget])
        dpg.add_button(label='1', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=add_digit, user_data=[entry_widget, 1])
        dpg.add_button(label='2', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=add_digit, user_data=[entry_widget, 2])
        dpg.add_button(label='3', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=add_digit, user_data=[entry_widget, 3])

    with dpg.group(horizontal=True):
        dpg.add_button(label='ctg(x)', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=make_ctg, user_data=[compute_widget, entry_widget, precision_widget])
        dpg.add_button(label='+/-', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=make_negative, user_data=[compute_widget, entry_widget, precision_widget])
        dpg.add_button(label='0', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=add_digit, user_data=[entry_widget, 0])
        dpg.add_button(label='.', width=BUTTON_WIDTH, height=BUTTON_HEIGHT,
                       callback=make_float, user_data=[compute_widget, entry_widget, precision_widget])


dpg.show_viewport()
dpg.set_primary_window('Primary', True)
dpg.start_dearpygui()
dpg.destroy_context()
