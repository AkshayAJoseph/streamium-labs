<script>
    import { onMount } from "svelte";
    let state = false;

    onMount(() => {
        const notch = document.querySelector("notch");
        const toggle = () => {
            state = !state;
        };
    });
</script>

<main>
    <div class="notch">
        <img src="/person-cv.svg" alt="" />
    </div>
    <div class="expand">
        <p class="rx">hey</p>
        <p class="tx">heyyy</p>
        <div class="expand__bottom">
            <input placeholder="Enter your query" type="text" />
            <input type="submit" {onclick} />
        </div>
    </div>
</main>

<style>
    .notch {
        position: fixed;
        bottom: 2rem;
        padding: 2rem;
        background: var(--color-white);
        color: var(--color-black);
        border-radius: 50%;
        height: 2vh;
        width: 2vh;
        display: flex;
        justify-content: center;
        align-items: center;
        right: 2rem;
    }
    .expand {
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        position: fixed;
        bottom: 8rem;
        background: var(--color-white);
        height: 40vh;
        width: 20vw;
        right: 2rem;
        padding: 2rem;
        border-radius: 2rem;
    }
    input {
        width: 100%;
        display: block;
        padding: 1rem 2rem;
        border: 0;
        border-radius: 1rem;
        background: #ddd;
        color: var(--color-black);
        outline: 0;
    }
    input[type="submit"] {
        background: var(--color-black);
        color: var(--color-white);
        border-radius: 5rem;
        margin-top: 2rem;
        font-weight: 600;
    }
    p.rx {
        background: var(--color-esports);
        border-radius: 2rem;
        padding: 1rem 2rem;
        width: fit-content;
        color: white;
    }
    p.tx {
        transform: translateX(180%);
        background: var(--color-esports);
        border-radius: 2rem;
        padding: 1rem 2rem;
        width: fit-content;
        color: white;
    }
</style>
