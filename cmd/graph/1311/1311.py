# 1311. Get Watched Videos by Your Friends

from queue import Queue

class Solution:
    def watchedVideosByFriends(self, watchedVideos: List[List[str]], friends: List[List[int]], id: int, level: int) -> List[str]:
        dict_freq = {}
        lev = 0
        fr_set = set()
        visited = [False]*len(friends)

        q = Queue()
        q.put(id)

        # find level friends
        while True:
            size = q.qsize()
            for i in range(size):
                fr_id = q.get()
                for fr in friends[fr_id]:
                    if not visited[fr]:
                        q.put(fr)
                        visited[fr] = True
            lev = lev + 1
            if level == lev:
                break

        # get unique friends id
        while not q.empty():
            fr = q.get()
            if fr != id:
                fr_set.add(fr)

        # count watched video frequency
        for fr in fr_set:
            for mov in watchedVideos[fr]:
                dict_freq[mov] = dict_freq.get(mov, 0) + 1

        # sort result
        sorts = sorted(dict_freq.items(), key=lambda item: (item[1], item[0]))

        # fill result
        return [k for k, v in sorts]