<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import Logo from '../../assets/Union.svg';

	/** @param {{ currentTarget: EventTarget & HTMLFormElement}} event */
	async function handleSubmit(event: { currentTarget: EventTarget & HTMLFormElement }) {
		const data = new FormData(event.currentTarget);

		if (data.get('email') === '') {
			/* TODO: error handling */
		}

		if (data.get('password') === '') {
			/* TODO: error handling */
		}

		data.set('login', data.get('email')?.toString().trim() as string);
		data.set('password', data.get('password')?.toString().trim() as string);

		const response = await fetch('http://localhost:8080/login', {
			method: 'POST',
			body: data
		});

		const result = deserialize(await response.text());

		if (result.type === 'success') {
			await invalidateAll();
			/* TODO: redirect to */
		}

		applyAction(result);
	}
</script>

<div class="screen-container">
	<header>
		<img src={Logo} alt="Logo" />
		<span><strong> GYM </strong></span>
	</header>

	<div class="container">
		<div class="form-container">
			<h1>Bem Vindo de Volta!</h1>
			<p>Log in para continuar administrando seus empregados.</p>
			<form method="POST" on:submit|preventDefault={handleSubmit}>
				<div class="form-input">
					<label for="email"> Email</label>
					<input name="email" type="email" />
				</div>
				<div class="form-input">
					<label for="password"> Senha </label>
					<input name="password" type="password" />
				</div>
				<button class="form-button">Login</button>
			</form>
			<p class="forgot-pass">Esqueceu sua senha?</p>
			<p class="help">Precisa de ajuda? <strong>Contate o suporte</strong></p>
		</div>
	</div>
</div>

<style>
	form {
		display: flex;
		flex-direction: column;
		cursor: pointer;
		gap: 1.5em;
	}

	.screen-container {
		width: 100vw;
		height: 100vh;

		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.container {
		width: 100vw;
		height: 70vw;

		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}

	.form-container {
		display: flex;
		flex-direction: column;
		text-align: center;

		width: 25%;
	}

	.form-input {
		display: flex;
		flex-direction: column;
	}

	.form-input > label {
		text-align: left;
	}

	.form-input > input {
		height: 3em;

		border-radius: 6px;
		border: 1px solid gray;
	}

	.form-button {
		background: black;
		color: white;

		cursor: pointer;

		height: 3em;

		font-weight: bold;
		font-size: 1em;

		border: none;
		border-radius: 6px;
	}

	.form-button:hover {
		opacity: 85%;
	}

	.forgot-pass {
		cursor: pointer;
	}

	.help {
		cursor: pointer;
	}

	header {
		display: flex;
		align-items: center;
		gap: 1em;

		margin-top: 1.5em;
	}

	input[type='email'] {
		outline-style: none;
		padding: 12px;
		margin: 8px 0;
		box-sizing: border-box;
	}

	input[type='password'] {
		outline-style: none;
		padding: 12px;
		margin: 8px 0;
		box-sizing: border-box;
	}
</style>
