def data_reader(text_file):
    file = open(text_file, "r")
    file_contents = [item.rstrip() for item in file]
    return file_contents