import matplotlib.pyplot as plt
import numpy as np

plt.style.use('_mpl-gallery')

x = []
y = []
with open('data.txt') as f:
    for line in f:
        data = line.split()
        x.append(int(data[0]) / 1000)
        y.append(float(data[1]))

plt.xscale('log')

plt.plot(x, y, linewidth=2)
plt.yticks(np.arange(0, max(y)+1, 5.0))
plt.ylim(ymin=0)
plt.xlabel("Data size, kilobytes")
plt.ylabel("Latency, nanoseconds")
plt.tight_layout()
plt.show()
