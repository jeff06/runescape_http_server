def complete_save(values: [], table: str):
    l = len(values)
    values[l - 1] = values[l - 1][:-1] + ";"

    with open(f'../database/query/{table}/insert.sql', 'w') as f:
        for line in values:
            f.write(f"{line}\n")


def fix_string_format(value: str):
    return value.replace("'", "''").replace("\n", "")


def fix_int_format(value: str):
    return value.replace("\n", "")