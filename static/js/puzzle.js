const end = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0]
class FifteenPuzzles {
    constructor(option) {
        this.board = []
        this.solution = ""
        this.startNode = []
        this.timer = option.animateTime || 1000
        this.init()

    }
    //根据当前状态渲染各数字位置
    calDom(node) {
        node.forEach((item, index) => {
            item.forEach((obj, i) => {
                $('#' + obj).css({ left: i * (100 + 2), top: index * (100 + 2) })
            })
        })
    }
    //每次状态改变调用一次渲染函数
    showDomMove(path) {
        let _ = this
        path.forEach(function (item, index) {
            setTimeout(function (node) {
                this.calDom(node)
            }.bind(_, item), index * _.timer)
        })
    }
    // 设置初始状态
    init() {
        let board = end
        this.board = board
        this.startNode = [
            [1, 2, 3, 4],
            [5, 6, 7, 8],
            [9, 10, 11, 12],
            [13, 14, 15, 0],
        ]
    }

    solve() {
        let _ = this
        let solution = this.solution.split("")
        let resultPath = getMove(solution, _.startNode)
        setTimeout(() => {
            _.showDomMove(resultPath)
        }, 500)
        this.board = end
    }
    setNode(board) {
        this.board = board
        this.startNode.forEach((item, index) => {
            item.forEach((obj, i) => {
                this.startNode[index][i] = board[index * 4 + i]
            })
        })
        // console.log(this.startNode)
        fifteenPuzzles.calDom(fifteenPuzzles.startNode)
    }
}
copy = (arr) => {
    let res = []
    arr.forEach(element => {
        res.push(element.concat())
    })
    return res
}

getZero = (node) => {
    node.forEach(function (item, i) {
        item.forEach(function (obj, j) {
            if (obj === 0) {
                target = { x: i, y: j }
            }
        })
    })
    return target
}
getMove = (solution, startNode) => {
    // console.log("sr=",startNode)
    zero = getZero(startNode)
    curr = copy(startNode)
    resultPath = [copy(curr)]
    action = {
        'U': [-1, 0],
        'D': [1, 0],
        'L': [0, -1],
        'R': [0, 1]
    }
    // console.log(resultPath)
    solution.forEach((element) => {
        next = {
            x: zero.x + action[element][0],
            y: zero.y + action[element][1]
        }
        temp = curr[zero.x][zero.y]
        curr[zero.x][zero.y] = curr[next.x][next.y]
        curr[next.x][next.y] = temp
        zero = next
        resultPath.push(copy(curr))
    })
    // console.log(resultPath)
    return resultPath
}

// 绑定按钮
bind = () => {
    document.getElementById('generate').onclick = getBoard
    document.getElementById('solution').onclick = getSolution
}

getBoard = () => {
    url = '/generate'
    var xhr = new XMLHttpRequest();
    xhr.open('get', url, true);
    xhr.send();
    xhr.onreadystatechange = function (resp) {
        if (xhr.readyState === 4) { // 读取完成
            if (xhr.status === 200) {
                resp = JSON.parse(xhr.responseText)
                board = resp['data']
                fifteenPuzzles.setNode(board)
            }
        }
    }
}
// 调整使其可解
adjust = () => {
    url = '/adjust'
    board = fifteenPuzzles.board
    var xhr = new XMLHttpRequest();
    xhr.open('post', url, true);
    xhr.setRequestHeader('content-type', 'application/json'); // 设置 HTTP 头，数据指定为 JSON 格式
    xhr.send(JSON.stringify({ "data": board }));
    xhr.onreadystatechange = function (resp) {
        if (xhr.readyState === 4) { // 读取完成
            if (xhr.status === 200)
                resp = JSON.parse(xhr.responseText)
            // 原来已经可解
            if (resp.code === 1004) {
                alert(resp.msg)
                return
            }
            fifteenPuzzles.setNode(resp.data)
        }
    }
}
// 获取结果
getSolution = () => {
    loading()
    url = '/solution'
    board = fifteenPuzzles.board
    var xhr = new XMLHttpRequest();
    xhr.open('post', url, true);
    xhr.setRequestHeader('content-type', 'application/json');
    xhr.send(JSON.stringify({ "data": board }));
    xhr.onreadystatechange = function (resp) {
        if (xhr.readyState === 4) { // 读取完成
            if (xhr.status === 200) {
                resp = JSON.parse(xhr.responseText)
                if (resp.code !== 1000) {
                    // 不可解
                    if (resp.code === 1002) {
                        if (confirm(resp.msg + ",是否调整？")) {
                            adjust()
                        }
                        loaded()
                        return
                    }
                    alert(resp.msg)
                    loaded()
                    return
                }
                loaded()
                alert(resp.msg + ",点击开始")
                fifteenPuzzles.solution = resp.data
                fifteenPuzzles.solve()
            }
        }
    }
}

loading = () => {
    document.getElementById('loading').style.display = "block"
}
loaded = () => {
    document.getElementById('loading').style.display = "none"

}