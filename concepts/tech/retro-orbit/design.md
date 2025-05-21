# Retro Orbit Design System

> A color system inspired by retro futurism and NASA punk aesthetics

RetroOrbit is a comprehensive color palette system designed for applications with a retro-futuristic and space exploration aesthetic. Inspired by the design philosophy of Catppuccin but reimagined through the lens of vintage NASA, retrofuturism, and space-age design.

## Design Philosophy

1. **Colorful is better than colorless**: Colors help distinguish between elements and make structure more clear
2. **There should be balance**: Not too dull, not too bright - colors work well in various lighting conditions
3. **Harmony is superior to dissonance**: Vibrant colors complement each other within a cohesive theme

## Palette Flavors

RetroOrbit offers four distinct palette flavors, each inspired by different eras and aesthetics of space exploration:

1. **Voyager**: A deep space-inspired dark theme with accent colors reminiscent of vintage NASA control panels and retro sci-fi designs
2. **Apollo**: A medium-dark theme inspired by the Apollo era space program with warm, slightly faded accent colors
3. **Launchpad**: A medium-light theme with colors inspired by technical documents, blueprints, and early NASA promotional materials
4. **Horizon**: A light theme inspired by Space Age optimism, retrofuturistic advertisements, and NASA's vision of space colonies

## Usage Guidelines

Each palette includes base colors for structural elements and accent colors with corresponding contrast colors:

- **Base Colors**: Use for backgrounds, surfaces, content areas, and text
- **Accent Colors**: Use for interactive elements, highlights, and decorative components
- **Contrast Colors**: Darker variants of accent colors for states (hover, pressed) and to ensure accessibility

## Implementation Guide

When implementing the RetroOrbit Design System:

1. **Select a flavor** that best matches your application's mood and target environment
2. **Use base colors** for structural elements (backgrounds, containers, text)
3. **Apply accent colors** for interactive elements, highlights, and decorative components
4. **Use contrast colors** for:
   - Hover/focus/active states
   - Borders and outlines
   - Text on accent-colored backgrounds
   - Ensuring accessible contrast ratios

## Accessibility

All color combinations meet or exceed WCAG 2.1 AA standards for text legibility:

- Normal text (smaller than 18pt): Minimum contrast ratio of 4.5:1
- Large text (larger than 18pt or bold): Minimum contrast ratio of 3:1

---

## Voyager (Dark Theme)

A deep space-inspired dark theme with accent colors reminiscent of vintage NASA control panels and retro sci-fi designs.

### Base Colors

| Color | Value | Preview |
|-------|-------|---------|
| Background | `#0A0E17` | ![Background](https://placehold.co/90x90/0A0E17/E6EDF3.png?text=Background) |
| Surface | `#141A24` | ![Surface](https://placehold.co/90x90/141A24/E6EDF3.png?text=Surface) |
| Overlay | `#1D2735` | ![Overlay](https://placehold.co/90x90/1D2735/E6EDF3.png?text=Overlay) |
| Text | `#E6EDF3` | ![Text](https://placehold.co/90x90/0A0E17/E6EDF3.png?text=Text) |
| Subtext | `#A0ABC0` | ![Subtext](https://placehold.co/90x90/0A0E17/A0ABC0.png?text=Subtext) |

### Accent & Contrast Colors

| Color | Accent Value | Contrast Value | Preview |
|-------|-------------|---------------|---------|
| Red | `#FF5A5F` | `#f30007` | ![Red](https://placehold.co/90x90/f30007/FF5A5F.png?text=Red) |
| Orange | `#FF9E64` | `#fd5f00` | ![Orange](https://placehold.co/90x90/fd5f00/FF9E64.png?text=Orange) |
| Yellow | `#FFCB6B` | `#ffa705` | ![Yellow](https://placehold.co/90x90/ffa705/FFCB6B.png?text=Yellow) |
| Green | `#7BD88F` | `#2ac34b` | ![Green](https://placehold.co/90x90/2ac34b/7BD88F.png?text=Green) |
| Teal | `#5CCCC6` | `#259d97` | ![Teal](https://placehold.co/90x90/259d97/5CCCC6.png?text=Teal) |
| Blue | `#69A7EB` | `#1072de` | ![Blue](https://placehold.co/90x90/1072de/69A7EB.png?text=Blue) |
| Purple | `#BF94E4` | `#8f37db` | ![Purple](https://placehold.co/90x90/8f37db/BF94E4.png?text=Purple) |
| Pink | `#F0A6CA` | `#ec4496` | ![Pink](https://placehold.co/90x90/ec4496/F0A6CA.png?text=Pink) |

### Accessibility Notes

- Text on Background: 16.33 (Excellent)
- Subtext on Background: 8.35 (Excellent)
- All accent colors meet or exceed 4.5:1 contrast ratio against the background when used with the contrast color variant

### Color Variables (CSS)

```css
/* Voyager Theme */
:root {
  /* Base */
  --voyager-bg: #0A0E17;
  --voyager-surface: #141A24;
  --voyager-overlay: #1D2735;
  --voyager-text: #E6EDF3;
  --voyager-subtext: #A0ABC0;
  
  /* Accent */
  --voyager-red: #FF5A5F;
  --voyager-orange: #FF9E64;
  --voyager-yellow: #FFCB6B;
  --voyager-green: #7BD88F;
  --voyager-teal: #5CCCC6;
  --voyager-blue: #69A7EB;
  --voyager-purple: #BF94E4;
  --voyager-pink: #F0A6CA;
  
  /* Contrast */
  --voyager-red-contrast: #f30007;
  --voyager-orange-contrast: #fd5f00;
  --voyager-yellow-contrast: #ffa705;
  --voyager-green-contrast: #2ac34b;
  --voyager-teal-contrast: #259d97;
  --voyager-blue-contrast: #1072de;
  --voyager-purple-contrast: #8f37db;
  --voyager-pink-contrast: #ec4496;
}
```

---

## Apollo (Medium Dark Theme)

A medium-dark theme inspired by the Apollo era space program with warm, slightly faded accent colors.

### Base Colors

| Color | Value | Preview |
|-------|-------|---------|
| Background | `#1A1D21` | ![Background](https://placehold.co/90x90/1A1D21/E8E3DC.png?text=Background) |
| Surface | `#25292E` | ![Surface](https://placehold.co/90x90/25292E/E8E3DC.png?text=Surface) |
| Overlay | `#303540` | ![Overlay](https://placehold.co/90x90/303540/E8E3DC.png?text=Overlay) |
| Text | `#E8E3DC` | ![Text](https://placehold.co/90x90/1A1D21/E8E3DC.png?text=Text) |
| Subtext | `#B6AEA1` | ![Subtext](https://placehold.co/90x90/1A1D21/B6AEA1.png?text=Subtext) |

### Accent & Contrast Colors

| Color | Accent Value | Contrast Value | Preview |
|-------|-------------|---------------|---------|
| Red | `#E35C53` | `#a61a10` | ![Red](https://placehold.co/90x90/a61a10/E35C53.png?text=Red) |
| Orange | `#E09356` | `#a35313` | ![Orange](https://placehold.co/90x90/a35313/E09356.png?text=Orange) |
| Yellow | `#E8C170` | `#c58b13` | ![Yellow](https://placehold.co/90x90/c58b13/E8C170.png?text=Yellow) |
| Green | `#65A373` | `#2c5c37` | ![Green](https://placehold.co/90x90/2c5c37/65A373.png?text=Green) |
| Teal | `#51A39A` | `#21544e` | ![Teal](https://placehold.co/90x90/21544e/51A39A.png?text=Teal) |
| Blue | `#5585B5` | `#234567` | ![Blue](https://placehold.co/90x90/234567/5585B5.png?text=Blue) |
| Purple | `#9A7AA0` | `#5b3962` | ![Purple](https://placehold.co/90x90/5b3962/9A7AA0.png?text=Purple) |
| Pink | `#CB829F` | `#9b325c` | ![Pink](https://placehold.co/90x90/9b325c/CB829F.png?text=Pink) |

### Accessibility Notes

- Text on Background: 13.25 (Excellent)
- Subtext on Background: 7.70 (Excellent)
- Many accent colors need their contrast color variants to meet accessibility requirements

### Color Variables (CSS)

```css
/* Apollo Theme */
:root {
  /* Base */
  --apollo-bg: #1A1D21;
  --apollo-surface: #25292E;
  --apollo-overlay: #303540;
  --apollo-text: #E8E3DC;
  --apollo-subtext: #B6AEA1;
  
  /* Accent */
  --apollo-red: #E35C53;
  --apollo-orange: #E09356;
  --apollo-yellow: #E8C170;
  --apollo-green: #65A373;
  --apollo-teal: #51A39A;
  --apollo-blue: #5585B5;
  --apollo-purple: #9A7AA0;
  --apollo-pink: #CB829F;
  
  /* Contrast */
  --apollo-red-contrast: #a61a10;
  --apollo-orange-contrast: #a35313;
  --apollo-yellow-contrast: #c58b13;
  --apollo-green-contrast: #2c5c37;
  --apollo-teal-contrast: #21544e;
  --apollo-blue-contrast: #234567;
  --apollo-purple-contrast: #5b3962;
  --apollo-pink-contrast: #9b325c;
}
```

---

## Launchpad (Medium Light Theme)

A medium-light theme with colors inspired by technical documents, blueprints, and early NASA promotional materials.

### Base Colors

| Color | Value | Preview |
|-------|-------|---------|
| Background | `#D5D8DB` | ![Background](https://placehold.co/90x90/D5D8DB/242D3A.png?text=Background) |
| Surface | `#C5CAD3` | ![Surface](https://placehold.co/90x90/C5CAD3/242D3A.png?text=Surface) |
| Overlay | `#B7BEC9` | ![Overlay](https://placehold.co/90x90/B7BEC9/242D3A.png?text=Overlay) |
| Text | `#242D3A` | ![Text](https://placehold.co/90x90/D5D8DB/242D3A.png?text=Text) |
| Subtext | `#536171` | ![Subtext](https://placehold.co/90x90/D5D8DB/536171.png?text=Subtext) |

### Accent & Contrast Colors

| Color | Accent Value | Contrast Value | Preview |
|-------|-------------|---------------|---------|
| Red | `#BD3B41` | `#4d1215` | ![Red](https://placehold.co/90x90/4d1215/BD3B41.png?text=Red) |
| Orange | `#C86A3F` | `#5b2a13` | ![Orange](https://placehold.co/90x90/5b2a13/C86A3F.png?text=Orange) |
| Yellow | `#D8914B` | `#784512` | ![Yellow](https://placehold.co/90x90/784512/D8914B.png?text=Yellow) |
| Green | `#457355` | `#0a150e` | ![Green](https://placehold.co/90x90/0a150e/457355.png?text=Green) |
| Teal | `#47807B` | `#0e201e` | ![Teal](https://placehold.co/90x90/0e201e/47807B.png?text=Teal) |
| Blue | `#3B6187` | `#0a141f` | ![Blue](https://placehold.co/90x90/0a141f/3B6187.png?text=Blue) |
| Purple | `#725A8C` | `#261a33` | ![Purple](https://placehold.co/90x90/261a33/725A8C.png?text=Purple) |
| Pink | `#A04E69` | `#3d1824` | ![Pink](https://placehold.co/90x90/3d1824/A04E69.png?text=Pink) |

### Accessibility Notes

- Text on Background: 9.71 (Excellent)
- Subtext on Background: 4.43 (Good for large text)
- Text colors have good contrast against most accent colors
- Contrast colors provide excellent contrast against the background

### Color Variables (CSS)

```css
/* Launchpad Theme */
:root {
  /* Base */
  --launchpad-bg: #D5D8DB;
  --launchpad-surface: #C5CAD3;
  --launchpad-overlay: #B7BEC9;
  --launchpad-text: #242D3A;
  --launchpad-subtext: #536171;
  
  /* Accent */
  --launchpad-red: #BD3B41;
  --launchpad-orange: #C86A3F;
  --launchpad-yellow: #D8914B;
  --launchpad-green: #457355;
  --launchpad-teal: #47807B;
  --launchpad-blue: #3B6187;
  --launchpad-purple: #725A8C;
  --launchpad-pink: #A04E69;
  
  /* Contrast */
  --launchpad-red-contrast: #4d1215;
  --launchpad-orange-contrast: #5b2a13;
  --launchpad-yellow-contrast: #784512;
  --launchpad-green-contrast: #0a150e;
  --launchpad-teal-contrast: #0e201e;
  --launchpad-blue-contrast: #0a141f;
  --launchpad-purple-contrast: #261a33;
  --launchpad-pink-contrast: #3d1824;
}
```

---

## Horizon (Light Theme)

A light theme inspired by Space Age optimism, retrofuturistic advertisements, and NASA's vision of space colonies.

### Base Colors

| Color | Value | Preview |
|-------|-------|---------|
| Background | `#EFF1F5` | ![Background](https://placehold.co/90x90/EFF1F5/1C2430.png?text=Background) |
| Surface | `#E1E5ED` | ![Surface](https://placehold.co/90x90/E1E5ED/1C2430.png?text=Surface) |
| Overlay | `#D1D7E3` | ![Overlay](https://placehold.co/90x90/D1D7E3/1C2430.png?text=Overlay) |
| Text | `#1C2430` | ![Text](https://placehold.co/90x90/EFF1F5/1C2430.png?text=Text) |
| Subtext | `#4C5A70` | ![Subtext](https://placehold.co/90x90/EFF1F5/4C5A70.png?text=Subtext) |

### Accent & Contrast Colors

| Color | Accent Value | Contrast Value | Preview |
|-------|-------------|---------------|---------|
| Red | `#CF3B43` | `#601115` | ![Red](https://placehold.co/90x90/601115/CF3B43.png?text=Red) |
| Orange | `#D7623B` | `#6b250e` | ![Orange](https://placehold.co/90x90/6b250e/D7623B.png?text=Orange) |
| Yellow | `#E19642` | `#7e480c` | ![Yellow](https://placehold.co/90x90/7e480c/E19642.png?text=Yellow) |
| Green | `#398060` | `#081811` | ![Green](https://placehold.co/90x90/081811/398060.png?text=Green) |
| Teal | `#2F8985` | `#061918` | ![Teal](https://placehold.co/90x90/061918/2F8985.png?text=Teal) |
| Blue | `#2C5F9B` | `#081626` | ![Blue](https://placehold.co/90x90/081626/2C5F9B.png?text=Blue) |
| Purple | `#6F5697` | `#261a3a` | ![Purple](https://placehold.co/90x90/261a3a/6F5697.png?text=Purple) |
| Pink | `#B34A6D` | `#4c1829` | ![Pink](https://placehold.co/90x90/4c1829/B34A6D.png?text=Pink) |

### Accessibility Notes

- Text on Background: 13.81 (Excellent)
- Subtext on Background: 6.18 (Excellent)
- Most accent colors have good contrast with the background
- Contrast colors provide maximum contrast for interactive elements

### Color Variables (CSS)

```css
/* Horizon Theme */
:root {
  /* Base */
  --horizon-bg: #EFF1F5;
  --horizon-surface: #E1E5ED;
  --horizon-overlay: #D1D7E3;
  --horizon-text: #1C2430;
  --horizon-subtext: #4C5A70;
  
  /* Accent */
  --horizon-red: #CF3B43;
  --horizon-orange: #D7623B;
  --horizon-yellow: #E19642;
  --horizon-green: #398060;
  --horizon-teal: #2F8985;
  --horizon-blue: #2C5F9B;
  --horizon-purple: #6F5697;
  --horizon-pink: #B34A6D;
  
  /* Contrast */
  --horizon-red-contrast: #601115;
  --horizon-orange-contrast: #6b250e;
  --horizon-yellow-contrast: #7e480c;
  --horizon-green-contrast: #081811;
  --horizon-teal-contrast: #061918;
  --horizon-blue-contrast: #081626;
  --horizon-purple-contrast: #261a3a;
  --horizon-pink-contrast: #4c1829;
}
```
