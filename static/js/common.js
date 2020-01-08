//IOS开关样式配置
var elem = document.querySelector('.js-switch');
var switchery = new Switchery(elem, {
    color: '#1AB394'
});
var config = {
    '.chosen-select': {},
}
for (var selector in config) {
    $(selector).chosen(config[selector]);
}