# With a line-separated text file input, return the number of measurements that are greater than the previous measurement.
def did_depth_increase(text):
    prev = 0
    count = -1
    lines = []
    with open(text) as f:
        lines = f.readlines()
    for depth in lines:
        if float(depth) > prev:
            count += 1
        prev = float(depth)
    return count


# With a line-separated text file input, return the number of 3-sum windows that are greater than the previous windows.
def three_measurement_window(text):
    count = -1
    prev = 0
    with open(text) as f:
        lines = f.readlines()
    for i in range(len(lines) - 2):
        if float(lines[i]) + float(lines[i+1]) + float(lines[i+2]) > prev:
            count += 1
        prev = float(lines[i]) + float(lines[i+1]) + float(lines[i+2])
    return count
