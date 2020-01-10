var common = {
    post(url, data = {}) {
        return new Promise((resolve, reject) => {
            $.ajax({
                url: url,
                method: "post",
                data: data,
                dataType: "json",
                success(res) {
                    console.log(res)
                    resolve(res)
                },
                fail(err) {
                    console.log(err)
                    reject(err)
                }
            })
        })
    },

    get(url, data = {}) {
        return new Promise((resolve, reject) => {
            $.ajax({
                url: url,
                method: "get",
                data: data,
                dataType: "json",
                success(res) {
                    console.log(res)
                    resolve(res)
                },
                fail(err) {
                    console.log(err)
                    reject(err)
                }
            })
        })
    },


    delete(url, data = {}) {
        return new Promise((resolve, reject) => {
            $.ajax({
                url: url,
                method: "delete",
                data: data,
                dataType: "json",
                success(res) {
                    console.log(res)
                    resolve(res)
                },
                fail(err) {
                    console.log(err)
                    reject(err)
                }
            })
        })
    },


    put(url, data = {}) {
        return new Promise((resolve, reject) => {
            $.ajax({
                url: url,
                method: "put",
                data: data,
                dataType: "json",
                success(res) {
                    console.log(res)
                    resolve(res)
                },
                fail(err) {
                    console.log(err)
                    reject(err)
                }
            })
        })
    },


    getList(url, params) {
        layui.use(['laytpl', 'laypage'], function () {
            var laytpl = layui.laytpl
            var laypage = layui.laypage

            laytpl.config({
                open: '{%',
                close: '%}'
            })

            $.ajax({
                url: url,
                method: "get",
                data: params,
                dataType: "json",
                success(res) {
                    console.log(res)
                    var data = res.data
                    var getTpl = tpl.innerHTML
                        ,view = document.getElementById('view');
                    laytpl(getTpl).render(data, function(html){
                        view.innerHTML = html;
                    });

                    laypage.render({
                        elem: 'AjaxPage',
                        count: data.count,
                        limit: data.per_page,
                        curr: params.page || 1,
                        jump: function (obj, first) {
                            if(!first) {
                                params.page = obj.curr
                                common.getList(url, params)
                            }
                        }
                    });
                }
            })
        })
    },

    getListNoPage (url, params = {}) {
        layui.use(['laytpl'], function () {
            var laytpl = layui.laytpl

            laytpl.config({
                open: '{%',
                close: '%}'
            })

            $.ajax({
                url: url,
                method: "get",
                data: params,
                dataType: "json",
                success(res) {
                    var data = res.data
                    var getTpl = tpl.innerHTML
                        ,view = document.getElementById('view');
                    laytpl(getTpl).render(data, function(html){
                        view.innerHTML = html;
                    });
                }
            })
        })
    },


    saveForm(url, elem, is_redirect = true) {
        console.log(is_redirect)
        var data = $(`#${elem}`).serialize()
        common.post(url, data).then(res => {
            layui.use(['layer'], function () {
                layer = layui.layer
                if(res.code == 1) {
                    layer.msg(res.msg, {
                        icon: 6,
                    })
                    if(is_redirect) {
                        setTimeout(function () {
                            window.history.go(-1)
                        }, 1500)
                    }else {
                        setTimeout(function () {
                            window.location.reload()
                        }, 1500)
                    }
                }else {
                    layer.msg(res.msg, {
                        icon: 5,
                        time: 1500
                    })
                }
            })
        })
    },

    del(url, id) {
        layui.use(['layer'], function () {
            layer.confirm("确定删除吗?", {
                icon: 3,
                title: "提示",
            }, function (index) {
                common.post(url, {id: id}).then(res => {
                    if(res.code == 1) {
                        layer.msg(res.msg, {icon: 6})
                    }else {
                        lsyer.msg(res.msg, {icon: 5})
                    }
                    setTimeout(function () {
                        window.location.reload()
                    }, 1500)
                })
                layer.close(index)
            })
        })
    }
}
