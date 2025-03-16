<script>
    import { onMount } from "svelte";
    import { goto } from "$app/navigation";

    let name = "";
    let email = "";
    let phone = "";
    let message = "";

    // Company Contact Info
    const companyEmail = "info@company.com";
    const companyPhone = "+123 456 7890";
    const companyLocation = { lat: 40.7128, lng: -74.0060 }; // Example coordinates (New York)

    // Handle Form Submission
    async function handleSubmit() {
        const formData = { name, email, phone, message };
        console.log("Form Data Submitted:", formData); // Replace with an API call to send data

        alert("Thank you! Your message has been sent.");
        goto("/"); // Redirect to homepage or thank-you page
    }
</script>

<h1>Contact Us</h1>

<section>
    <h2>Send Us a Message</h2>
    <form on:submit|preventDefault={handleSubmit}>
        <label>
            Name:
            <input type="text" bind:value={name} required>
        </label>

        <label>
            Email:
            <input type="email" bind:value={email} required>
        </label>

        <label>
            Phone:
            <input type="tel" bind:value={phone} required>
        </label>

        <label>
            Message:
            <textarea bind:value={message} required></textarea>
        </label>

        <button type="submit">Submit</button>
    </form>
</section>

<section>
    <h2>Contact Information</h2>
    <p><strong>Email:</strong> {companyEmail}</p>
    <p><strong>Phone:</strong> {companyPhone}</p>
</section>

<section>
    <h2>Our Location</h2>
    <iframe
        width="100%"
        height="300"
        frameborder="0"
        style="border:0"
        src="https://www.google.com/maps?q={companyLocation.lat},{companyLocation.lng}&output=embed"
        allowfullscreen>
    </iframe>
</section>
