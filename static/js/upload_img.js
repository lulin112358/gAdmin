function wp_upload_img(fileList,server,pick,image_src,image_input) {
    var $list = fileList;
    //上传图片,初始化WebUploader
    var uploader = WebUploader.create({
        auto: true,// 选完文件后，是否自动上传。
        swf: '/static/admin/js/webupload/Uploader.swf',// swf文件路径
        server: server,// 文件接收服务端。
        duplicate :true,// 重复上传图片，true为可重复false为不可重复
        pick: pick,// 选择文件的按钮。可选。
        accept: {
            title: 'Images',
            extensions: 'gif,jpg,jpeg,bmp,png',
            mimeTypes: 'image/jpg,image/jpeg,image/png'
        }
    });
    // 文件上传成功
    uploader.on( 'uploadSuccess', function( file, data, response ) {
        if(data.code==1)
        {
            $(image_src).val(data.data);
            $(image_input).attr('src', data.data).show();
            layer.msg(data.msg, {icon: 6,time:1500,shade: 0.1}, function(index){
                layer.close(index);
            });
        }else{
            layer.msg(data.msg, {icon: 5,time:1500,shade: 0.1}, function(index){
                layer.close(index);
            });
        }
    });
    // 文件上传失败，显示上传出错。
    uploader.on( 'uploadError', function(file) {
        layer.msg('上传错误', {icon: 5,time:1500,shade: 0.1}, function(index){
            layer.close(index);
        });
    });
}


