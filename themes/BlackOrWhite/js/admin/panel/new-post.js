// 假设这是你的标签数组
const tags = [];

function renderTags() {
    // 获取ul元素
    const tagSelect = document.getElementById('TagSelect');

    // 清空ul中的内容以避免重复
    tagSelect.innerHTML = '';

    // 遍历tags数组，为每个标签创建一个li元素
    tags.forEach((tag, index) => {
        // 创建li元素
        const li = document.createElement('li');
        li.classList.add('list-group-item', 'd-flex', 'justify-content-between', 'align-items-center');

        // 设置li元素的文本内容为当前标签名称
        li.textContent = tag;

        // 创建关闭按钮
        const button = document.createElement('button');
        button.setAttribute('type', 'button');
        button.classList.add('btn-close');
        button.setAttribute('aria-label', 'Close');
        button.addEventListener('click', function() {
            // 从数组中删除当前标签
            tags.splice(index, 1);
            // 重新渲染标签
            renderTags();
        });

        // 将关闭按钮添加到li元素中
        li.appendChild(button);

        // 将li元素添加到ul中
        tagSelect.appendChild(li);
    });
}

document.getElementById('inputLabel').addEventListener('keydown', function(event) {
    if (event.key === 'Enter') {
        const newTag = this.value.trim();
        if (newTag) {
            tags.push(newTag);
            renderTags();
            this.value = ''; // 清空输入框
        }
    }
});

async function submitPost() {
    // 获取文章标题和内容
    const title = document.getElementById('postTitle').value;
    const content = simpleMde.value();

    // 获取选中的分类
    const category = document.querySelector('input[name="listGroupRadio"]:checked').nextElementSibling.textContent;

    // 准备要发送的数据
    const postData = {
        title,
        content,
        category,
        tags
    };

    try {
        const response = await fetch('/api/post', {
            method: 'POST', // 指定请求方法为POST
            headers: {
                'Content-Type': 'application/json', // 设置请求头，指明发送的数据类型为JSON
            },
            body: JSON.stringify(postData) // 将JavaScript对象转换为JSON字符串
        });

        if (response.ok) {
            // 请求成功，可以根据需要处理响应
            const responseData = await response.json(); // 假设服务器返回的是JSON数据
            console.log('Post submitted successfully:', responseData);
            alert('文章发布成功！');
        } else {
            // 服务器返回了错误状态码
            console.error('Failed to submit post:', response.status);
            alert('文章发布失败，请稍后重试。');
        }
    } catch (error) {
        // 网络错误或无法发送请求
        console.error('Error submitting post:', error);
        alert('提交时出现错误，请检查网络连接并重试。');
    }
}

// 给发布按钮添加点击事件监听器
document.querySelector('button.btn-outline-dark').addEventListener('click', submitPost);

// 首次渲染标签列表
renderTags();
let simpleMde = new SimpleMDE({ element: document.getElementById("postContent") });