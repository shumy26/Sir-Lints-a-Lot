from structures import Time, TimeBlock, DayBlock, Week

def find_index_of_timeblock(day, time):
    for idx, block in enumerate(day.times):
        if block.get_start_time() == time:
            return idx
    raise ValueError("Time not in day!")

def find_index_of_activity(day, activity):
    for idx, block in enumerate(day.times):
        if block.activity == activity:
            return idx
    raise Exception("Activity not in day!")

def match_activity_set_time(day, activity):
    amount_of_blocks_needed = activity.duration / Time(0,30)
    is_set_time = True if activity.set_time else False

    if is_set_time:
        start_time_idx = find_index_of_timeblock(day, activity.set_time)
        if start_time_idx + amount_of_blocks_needed <= len(day.times):
            if all(day.times[start_time_idx + j].is_free for j in range(amount_of_blocks_needed)): 
                for i in range(amount_of_blocks_needed):
                    day.times[start_time_idx + i].add_activity(activity)
                print(f"Added at {TimeBlock(day.times[start_time_idx].get_start_time(), activity.duration, activity.energy, activity=activity)}" )
                return True, day.times[start_time_idx].get_start_time()
        print("failed to allocate to set time, but will try at other times")
    return False, None

def match_activity(day, activity):
    set_time_successful = match_activity_set_time(day, activity)[0]
    if set_time_successful:
        return True, activity.set_time
    amount_of_blocks_needed = activity.duration / Time(0,30)
    for index in range(len(day.times)):
        if index + amount_of_blocks_needed > len(day.times):
            break
        if all(day.times[index + i].is_free and day.times[index + i].get_energy() == activity.energy \
                for i in range(amount_of_blocks_needed)):
            for i in range(amount_of_blocks_needed):
                day.times[index + i].add_activity(activity)
            start = day.times[index].get_start_time()
            print(f"Added at {TimeBlock(start, activity.duration, activity.energy, activity=activity)}" )
            return True, day.times[index].get_start_time()
    print("failed")    
    return False, None

def remove_activity_from_time(day, time):
    time_idx = find_index_of_timeblock(day, time)
    if day.times[time_idx].is_free:
        raise Exception("The time was already free!")
    activity = day.times[time_idx].activity
    while time_idx < len(day.times) and day.times[time_idx].activity == activity:
        day.times[time_idx].activity = None
        day.times[time_idx].is_free = True
        time_idx += 1
    
def remove_whole_activity(day, activity_name):
    check = False
    for time in day.times:
        if not time.is_free and time.activity.name.lower() == activity_name.lower():
            activity = time.activity
            check = True
            break
    if check:
        start_time = day.times[find_index_of_activity(day, activity)].get_start_time()
        remove_activity_from_time(day, start_time)
        print(f"Removed activity {activity_name} at time {start_time}")
    else:
        print(f"Could not remove {activity_name}")