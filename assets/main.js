(async () => {
    const counter = document.getElementsByClassName("js-counter")[0];
    const {
        num
    } = await fetch("/counter.json").then(resp => resp.json());
    counter.innerHTML = num;

    const incrementCounter = async () => {
        const num = parseInt(counter.innerHTML, 10) + 1;
        await fetch("/counter", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                num
            }),
        });
        counter.innerHTML = num;
    };

    const increment = document.getElementsByClassName("js-increment")[0];
    increment.addEventListener("click", incrementCounter)
})();
