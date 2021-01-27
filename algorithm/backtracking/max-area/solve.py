# -*- coding: utf-8 -*-

def main(argv):
    oMap = CreateMap(argv)
    dMap = CloneMap(oMap)
    maxPart = MarkMap(dMap)
    return FilterTestPoint(oMap, dMap, maxPart)


def CloneMap(oMap):
    dMap = []
    for i in range(oMap.m_iWidth):
        ith = []
        for j in range(oMap.m_iHeight):
            ith.append(oMap.m_MapData[i][j])
        dMap.append(ith)
    return dMap


def MarkMap(dMap):
    """
    将原来的0替换为具体的标记
    :param dMap:
    :return:
    """
    no = 2
    maxCount = 0
    maxPart = no
    for i in range(len(dMap)):
        for j in range(len(dMap[0])):
            if dMap[i][j] == 0:
                count = MarkPart(dMap, no, i, j)
                if count > maxCount:
                    maxCount = count
                    maxPart = no
    return maxPart


def MarkPart(dMap, no, i, j):
    """
    按照区域进行划分
    :param dMap:
    :param no:
    :param i:
    :param j:
    :return:
    """
    stack = [(i, j)]
    count = 0
    while len(stack) > 0:
        point = stack.pop()
        count += 1
        dMap[point[0]][point[1]] = no
        directions = [[1, 0], [-1, 0], [0, 1], [0, -1]]
        for direction in directions:
            newPoint = (point[0] + direction[0], point[1] + direction[1])
            if (
                newPoint[0] < 0
                or newPoint[0] >= len(dMap)
                or newPoint[1] < 0
                or newPoint[1] >= len(dMap[0])
            ):
                continue
            if dMap[newPoint[0]][newPoint[1]] != 0:
                continue
            stack.append(newPoint)
    return count


def GetMaxPart(countMap):
    """
    获取最大区域的编号
    :param countMap:
    :return:
    """
    maxValue = 0
    maxPart = 2
    for key in countMap.keys():
        if countMap[key] > maxValue:
            maxPart = key
            maxValue = countMap[key]
    return maxPart


def FilterTestPoint(oMap, dMap, maxPart):
    """
    过滤出落在最大区域的点
    :param oMap:
    :param dMap:
    :param maxPart:
    :return:
    """
    filtered = []
    for i in oMap.m_TestPoint:
        if (
            0 <= i[0] < len(dMap)
            and 0 <= i[1] < len(dMap[0])
            and dMap[i[0]][i[1]] == maxPart
        ):
            filtered.append(i)
    return filtered


def CreateMap(dInfo):
    return CMap.CreateMap(dInfo)


class CMap(object):
    def __init__(self):
        self.m_iWidth = 0
        self.m_iHeight = 0
        self.m_MapData = []
        self.m_TestPoint = []

    def SetData(self, mapData):
        """
        m_MapData[x][y]为坐标(x,y)的阻挡信息
        :param mapData:
        :return:
        """
        self.m_MapData = mapData
        self.m_iWidth = len(mapData)  # 地图宽
        self.m_iHeight = len(mapData[0])  # 地图高

    def IsBlock(self, x, y):
        """
        采用笛卡尔坐标系
        :param x:横坐标
        :param y:纵坐标
        :return:0:表示非阻挡，1表示阻挡
        """
        return self.m_MapData[x][y]

    def SetTestPoint(self, testPoint):
        """
        用于测试是否在最大闭合区域的点
        :param testPoint:
        :return:
        """
        self.m_TestPoint = map(tuple, testPoint)

    @classmethod
    def CreateMap(cls, dInfo):
        """
        地图数据的加载
        :param fpath:
        """
        mapData = dInfo.get("MapData", [])  # 获得地图信息
        testPoint = dInfo.get("TestPoint", [])  # 获得测试点数据
        oMap = CMap()
        oMap.SetData(mapData)
        oMap.SetTestPoint(testPoint)
        return oMap


if __name__ == "__main__":
    argv = {
        "MapData": (
            (1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1),
            (1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1),
            (1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1),
            (1, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 1, 1),
            (1, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 1, 1),
            (1, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1),
            (1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 0, 1, 1),
            (1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1),
            (1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1),
            (1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1),
            (1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1),
            (1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1),
            (1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1),
            (1, 0, 0, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1),
            (1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1),
            (1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 1),
            (1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1),
            (1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1),
            (1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1),
            (1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1),
            (1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1),
        ),
        "TestPoint": ((16, 15), (18, 4), (13, 2), (5, 2), (4, 12), (14, 2)),
    }
    result = main(argv)
    print(result)
