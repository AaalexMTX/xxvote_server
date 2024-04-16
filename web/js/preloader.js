document.onreadystatechange = function () {
    if (document.readyState === 'complete') {
        setTimeout(function () {
            // 页面已完全加载并且已经等待了至少一秒
            $(".hidden").removeClass("hidden");
            $(".loading").hide();
        }, 1000); // 1000毫秒（1秒）的延时
    }
};