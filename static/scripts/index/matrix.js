function getMarginRatio() {
    let marginPercent = style.getPropertyValue('--margin-ratio').trim();
    let marginValue = parseInt(marginPercent.slice(0, -1));
    return (marginValue / 100);
}

/* Credits: https://dev.to/gnsp/making-the-matrix-effect-in-javascript-din
   (with some modification) */
const canvas = document.getElementById('matrix-canvas');
const ctx = canvas.getContext('2d');

let w = canvas.width = document.body.offsetWidth;
let h = canvas.height = viewport.height;
let cols = Math.floor(w / 20) + 1;
let ypos = Array(cols).fill(h + 10);

let marginPercent = getMarginRatio();
let foregroundColor = style.getPropertyValue('--color-foreground-1000').trim();
let backgroundColor = style.getPropertyValue('--color-background-1000').trim();

function changeMatrixTheme(theme) {
    foregroundColor = style.getPropertyValue('--color-foreground-1000').trim();
    backgroundColor = style.getPropertyValue('--color-background-1000').trim();
    ctx.fillStyle = backgroundColor;
    ctx.fillRect(0, 0, w, h);
}

function updateSize() {
    const viewport = getViewport();
    w = canvas.width = document.body.offsetWidth;
    h = canvas.height = viewport.height;

    const newCols = Math.floor(w / 20) + 1;
    marginPercent = getMarginRatio();

    ctx.fillStyle = backgroundColor;
    ctx.fillRect(0, 0, w, h);

    if (newCols > cols) {
        ypos = ypos.concat(Array(newCols - cols + 2).fill(h + 10));
    } else if (newCols < cols) {
        ypos = ypos.slice(0, newCols);
    }
    cols = newCols;
    marginCols = cols * marginPercent;
}

if (window.changeThemeCallbacks !== undefined) {
    window.changeThemeCallbacks.push(changeMatrixTheme)
} else {
    window.changeThemeCallbacks = [changeMatrixTheme];
}

const debouncedUpdateSize = debounce(updateSize, 500);
window.addEventListener('resize', debouncedUpdateSize);

ctx.fillStyle = backgroundColor;
ctx.fillRect(0, 0, w, h);

function matrix() {
    ctx.fillStyle = backgroundColor + "33";  // adding alpha
    ctx.fillRect(0, 0, w, h);

    ctx.font = '10pt monospace';

    let marginCols = cols * marginPercent;

    ypos.forEach((y, ind) => {
        if (ind > marginCols && ind < cols - marginCols) {
            ctx.fillStyle = foregroundColor + "22"; // adding alpha
        } else {
            ctx.fillStyle = foregroundColor;
        }
        const text = String.fromCharCode(Math.random() * 128);
        const x = ind * 20;
        ctx.fillText(text, x, y);
        if (y > 100 + Math.random() * 10000) ypos[ind] = 0;
        else ypos[ind] = y + 20;
    });
}

setInterval(matrix, 100);
