function findMajorityElement(arr):
    candidate := -1
    count := 0
    for i := 0 to length(arr) - 1 do
        if count = 0 then
            candidate := arr[i]
            count := 1
        else if candidate = arr[i] then
            count := count + 1
        else
            count := count - 1
    count := 0
    for i := 0 to length(arr) - 1 do
        if candidate = arr[i] then
            count := count + 1

    if count > length(arr) / 2 then
        return candidate
    else
        return -1
