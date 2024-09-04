<script lang="ts">
    import { onMount } from 'svelte';

    let token: string = '';

    // APIリクエストを実行する関数
    async function fetchAccessToken() {
        try {
            const response = await fetch('http://localhost:5050/api/v1/auth/access_token', {
                method: 'GET', // curlではGETリクエストが使用されています
                headers: {
                    'x-user-name': 'admin',
                    'x-password': 'password'
                }
            });

            if (response.ok) {
                const data = await response.json();

                console.log(JSON.stringify(data));

                // token = JSON.stringify(data);

                token = data.data.access_token;

                // アクセストークンをlocalStorageに保存
                localStorage.setItem('access_token', token);
                console.log('Access token:', token);
            } else {
                console.error('Failed to fetch access token', response.status);
            }
        } catch (error) {
            console.error('Error:', error);
        }
    }

    // コンポーネントがマウントされたらAPIリクエストを送信
    onMount(() => {
        fetchAccessToken();
    });
</script>

<div>
    <h1>アクセストークン: {token}</h1>
</div>

