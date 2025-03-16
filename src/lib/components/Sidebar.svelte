<script>
    export let selected = "";
    import { logout } from "../../script";
</script>

<main>
    <aside>
        <div class="links">
            <div class="links__top">
                <h5>Services</h5>
                <a href="/dashboard/live">
                    <div class="links__row" class:super={selected === "live"}>
                        <img src="/live-alt.svg" alt="icon" />
                        <p>Live Stream</p>
                    </div>
                </a>
                <a href="/dashboard/media">
                    <div class="links__row" class:super={selected === "media"}>
                        <img src="/heart.svg" alt="icon" />
                        <p>Media Production</p>
                    </div>
                </a>
                <a href="/dashboard/events">
                    <div class="links__row" class:super={selected === "events"}>
                        <img src="/calendar-star.svg" alt="icon" />
                        <p>Event Management</p>
                    </div>
                </a>

                <h5>Miscellaneous</h5>
                <a href="/dashboard/profile">
                    <div
                        class="links__row"
                        class:super={selected === "profile"}
                    >
                        <img src="/user-trust.svg" alt="icon" />
                        <p>Profile Management</p>
                    </div>
                </a>

                <a href="/dashboard/hiring">
                    <div class="links__row" class:super={selected === "hiring"}>
                        <img src="/person-cv.svg" alt="icon" />
                        <p>Apply for Jobs</p>
                    </div>
                </a>
            </div>
            <div class="links__bottom">
                <a onclick={logout}>
                    <div class="links__row end">
                        <img src="/github.svg" alt="icon" />
                        <p>Logout</p>
                    </div>
                </a>
            </div>
        </div>
    </aside>
</main>

<style>
    aside {
        background: #111;
        width: 25vw;
        height: 100vh;
    }
    section {
        width: 60vw;
    }

    .links {
        padding: 1rem 2rem;
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        height: 100%;
    }

    .links h5 {
        font-size: 1.5rem;
        padding: 1rem;
    }

    .links__row {
        display: flex;
        gap: 1rem;
        align-items: center;
        padding: 1rem;
        border-radius: 1rem;
    }

    .links__row.super {
        background: #66666633;
        color: white;
    }

    .links__row p {
        transition: color ease 0.3s;
    }

    .links__row img {
        transition: filter ease 0.3s;
        width: 1rem;
        filter: invert(35%) sepia(59%) saturate(3%) hue-rotate(28deg)
            brightness(100%) contrast(97%);
    }

    .links__row:hover p {
        color: white;
    }

    .links__row:hover img {
        filter: invert(100%) sepia(0%) saturate(0%) hue-rotate(93deg)
            brightness(103%) contrast(103%);
    }
    @media only screen and (max-width: 640px) {
        aside {
            width: 100%;
            height: fit-content;
        }
    }
</style>
