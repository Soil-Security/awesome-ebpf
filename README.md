# Awesome Berkeley Packet Filter (BPF)

A very opinionated set of awesome resources that I read or bumped into while leraning eBPF.

## Fundamentals

* [Steven McCanne, Van Jacobson, December 1992, The BSD Packet Filter: A New Architecture for User-level Packet Capture](https://www.tcpdump.org/papers/bpf-usenix93.pdf)
* [Jonathan Corbet, April 2011, LWN.net, A JIT for packet filters](https://lwn.net/Articles/437981/)
* [Jonathan Corbet, January 2012, LWN.net, Yet another new approach to seccomp](https://lwn.net/Articles/475043/)
* [Alexei Starovoitov, 30 September 2013, LKML.org, extended BPF](https://lkml.org/lkml/2013/9/30/627)
* [Jonathan Corbet, June 2021, LWN.net, Spectre revisits BPF](https://lwn.net/Articles/860597/)
* [Alexei Starovoitov, August 2022, Kernel Recipes 2022 - The untold story of BPF](https://www.youtube.com/watch?v=DAvZH13725I)
* [John Fastabend, December 2022, Isovalent, How the Hive Came to Bee - The History of eBPF](https://www.youtube.com/watch?v=sYxUOKi3Q00)
* [Suchakra Sharma, IO Visor Documentation, BPF Internals - I](https://github.com/iovisor/bpf-docs/blob/master/bpf-internals-1.md)
* [Suchakra Sharma, IO Visor Documentation, BPF Internals - II](https://github.com/iovisor/bpf-docs/blob/master/bpf-internals-2.md)
* [Kernel Documentation, BPF Documentation](https://www.kernel.org/doc/html/latest/bpf/)
* [Kernel Documentation, BPF Type Format (BTF)](https://www.kernel.org/doc/Documentation/bpf/btf.rst)
* [Kernel Documentation, Linux Socket Filtering aka Berkeley Packet Filter (BPF)](https://www.kernel.org/doc/Documentation/networking/filter.txt)
* [Kernel Documentation, BPF Design QA](https://www.kernel.org/doc/Documentation/bpf/bpf_design_QA.txt)
* [Kernel Documentation, Using the Linux Kernel Tracepoints](https://www.kernel.org/doc/Documentation/trace/tracepoints.rst)
* [Linux manual page, bpf(2)](https://man7.org/linux/man-pages/man2/bpf.2.html)
* [Linux manual page, tc-bpf(8)](https://man7.org/linux/man-pages/man8/tc-bpf.8.html)
* [Linux manual page, bpf-helpers(7)](https://man7.org/linux/man-pages/man7/bpf-helpers.7.html)
* [Cilium Documentation: BPF and XDP Reference Guide](https://docs.cilium.io/en/v1.13/bpf/index.html)
* [Cilium Documentation: eBPF Datapath](https://docs.cilium.io/en/v1.13/network/ebpf/#ebpf-datapath)
* [IO Visor Documentation, Unofficial eBPF spec](https://github.com/iovisor/bpf-docs/blob/master/eBPF.md)
* [IO Visor Documentation, BPF Features by Linux Kernel Version](https://github.com/iovisor/bcc/blob/master/docs/kernel-versions.md)
* [Wikipedia, Berkeley Packet Filter](https://en.wikipedia.org/wiki/Berkeley_Packet_Filter)

## Arcana

* [Sudhanshu Goswami, April 2005, LWN.net, An introduction to KProbes](https://lwn.net/Articles/132196/)
* [Jonathan Corbet, May 2014, LWN.net, BPF: the universal in-kernel virtual machine](https://lwn.net/Articles/599755/)
* [Jonathan Corbet, July 2014, LWN.net, Extending extended BPF](https://lwn.net/Articles/603983/)
* [Jonathan Corbet, September 2014, LWN.net, The BPF system call API, version 14](https://lwn.net/Articles/612878/)
* [Matt Fleming, December 2017, LWN.net, A thorough introduction to eBPF](https://lwn.net/Articles/740157/)
* [Matt Fleming, May 2018, LWN.net, Using user-space tracepoints with BPF](https://lwn.net/Articles/753601/)
* [Jonathan Corbet, December 2018, LWN.net, Bounded loops in BPF programs](https://lwn.net/Articles/773605/)
* [Jake Edge, September 2019, LWN.net, Kernel runtime security instrumentation](https://lwn.net/Articles/798157/)
* [Nurkholish Halim, September 2022, Medium, A Deep Dive into eBPF: Writing an Efficient DNS Monitoring](https://medium.com/@nurkholish.halim/a-deep-dive-into-ebpf-writing-an-efficient-dns-monitoring-2c9dea92abdf)
* [Adrian Ratiu, April 2019, Collabora, An eBPF overview, part 1: Introduction](https://www.collabora.com/news-and-blog/blog/2019/04/05/an-ebpf-overview-part-1-introduction/)
* [Adrian Ratiu, April 2019, Collabora, An eBPF overview, part 2: Machine & bytecode](https://www.collabora.com/news-and-blog/blog/2019/04/15/an-ebpf-overview-part-2-machine-and-bytecode/)
* [Securing Linux with a Faster and Scalable IPtables](https://ccronline.sigcomm.org/wp-content/uploads/2019/07/acmdl19-304.pdf)
* [Luuk Hendriks, July 2020, RIPE Labs, Journeying into XDP: Part 0](https://labs.ripe.net/author/luuk_hendriks/journeying-into-xdp-part-0/)
* [Tom Carpay, October 2020, RIPE Labs, Journeying into XDP Part 1: Augmenting DNS](https://labs.ripe.net/author/tom_carpay/journeying-into-xdp-part-1-augmenting-dns/)
* [Luuk Hendriks, July 2020, RIPE Labs, Journeying into XDP Part 2: XDPerimenting with DNS Telemetry](https://labs.ripe.net/author/luuk_hendriks/journeying-into-xdp-part-2-xdperimenting-with-dns-telemetry/)
* [Sanjeev Rampal, Donald Hunter, May 2023, eBPF 201: Supercharging Your eBPF Dev Process for Cloud Native Apps](https://www.youtube.com/watch?v=9vUJhbqf4eo)

## XDP

* [October 2018, Load XDP programs using the ip (iproute2) command](https://fntlnz.wtf/post/xdp-ip-iproute/)
* [Kong Blog, October 2022, Writing an eBPF/XDP load-balancer in Rust](https://konghq.com/blog/engineering/writing-an-ebpf-xdp-load-balancer-in-rust)

## Computer Science

* [Wikipedia, Control-flow graph](https://en.wikipedia.org/wiki/Control-flow_graph)
* [Yunhe Shi, David Gregg, Andrew Beatty, M. Anton Ertl, Virtual Machine Showdown: Stack Versus Registers](https://www.usenix.org/legacy/events/vee05/full_papers/p153-yunhe.pdf)
* [Hongjiu Lu, May 1995, ELF: From The Programmer’s Perspective](https://www.linux.co.cr/free-unix-os/review/acrobat/950517.pdf)
* [June 2019, Simple and Precise Static Analysis of Untrusted Linux Kernel Extensions](https://dl.acm.org/doi/pdf/10.1145/3314221.3314590)
* [The eXpress data path: fast programmable packet processing in the operating system kernel](https://dl.acm.org/doi/10.1145/3281411.3281443)

## Books

* [Liz Rice, March 2023, O'Reilly, Learning eBPF](https://www.oreilly.com/library/view/learning-ebpf/9781098135119/)
* [David Calavera, Lorenzo Fontana, September 2019, Linux Observability with BPF](https://www.oreilly.com/library/view/linux-observability-with/9781492050193/)
* [Brendan Gregg, December 2019, Addison-Wesley Professional, BPF Performance Tools](https://www.oreilly.com/library/view/bpf-performance-tools/9780136588870/)
* [Brendan Gregg, November 2020, Addison-Wesley, Systems Performance: Enterprise and the Cloud](https://www.brendangregg.com/blog/2020-07-15/systems-performance-2nd-edition.html)
* [Jed Salazar, Natalia Reka Ivanko, April 2022, O'Reilly, Security Observability with eBPF](https://www.oreilly.com/library/view/security-observability-with/9781492096719/)
* [Liz Rice, April 2022, O'Reilly, What is eBPF?](https://www.oreilly.com/library/view/what-is-ebpf/9781492097266/)

## Compilers

* [Quentin Monnet, April 2020, eBPF assembly with LLVM](https://qmonnet.github.io/whirl-offload/2020/04/12/llvm-ebpf-asm/)
* [Jonathan Corbet, September 2019, LWN.net, Compiling to BPF with GCC](https://lwn.net/Articles/800606/)
* [Jake Edge, September 2020, LWN.net, BPF in GCC](https://lwn.net/Articles/831402/)

## Code Examples

* [:lock: danielpacak/bpf-c-template](https://github.com/danielpacak/bpf-c-template/)
* [:lock: danielpacak/bpf-go-template](https://github.com/danielpacak/bpf-go-template/)
* [:lock: danielpacak/bpf-xdp-go-template](https://github.com/danielpacak/bpf-xdp-go-template/)
* [:lock: danielpacak/bpf-xdp-rust-template](https://github.com/danielpacak/bpf-xdp-rust-template/)
* [:lock: danielpacak/bpf-tc-go-template](https://github.com/danielpacak/bpf-tc-go-template/)
* [:lock: danielpacak/bpf-windows-template](https://github.com/danielpacak/bpf-windows-template/)
* [libbpf-bootstrap: demo BPF applications](https://github.com/libbpf/libbpf-bootstrap/)
* [David Calavera, Lorenzo Fontana, BPF Workshop](https://github.com/bpftools/bpf-workshop/)
* [xdp-project/xdp-tutorial](https://github.com/xdp-project/xdp-tutorial/)
* [NLnetLabs/XDPeriments](https://github.com/NLnetLabs/XDPeriments/)
* [eunomia-bpf/bpf-developer-tutorial](https://github.com/eunomia-bpf/bpf-developer-tutorial/)

## Tools and Frameworks

* [Quentin Monnet, September 2021, Features of bpftool: the thread of tips and examples to work with eBPF objects](https://qmonnet.github.io/whirl-offload/2021/09/23/bpftool-features-thread/)
* [Michael Kerrisk, February 2020, Linux Security and Isolation APIs Seccomp](https://man7.org/training/download/secisol_seccomp_slides.pdf)
* [edgebitio/edgebit-agent](https://github.com/edgebitio/edgebit-agent)
* [:lock: danielpacak/lazybpftool](https://github.com/danielpacak/lazybpftool)
* [bpfd-dev/bpfd](https://github.com/bpfd-dev/bpfd)
* [rust-bpf/rust-bcc](https://github.com/rust-bpf/rust-bcc)
* [October 2021, eCHO episode 25: eBPF, Rust and Aya](https://www.youtube.com/watch?v=TQ0ou-eFLAk)
* [Quentin Monnet, February 2020, Tools and mechanisms to debug BPF programs](https://archive.fosdem.org/2020/schedule/event/debugging_bpf/)
* [Udi Rot, January 2023, Navigate your way to production bliss with Caretta](https://www.groundcover.com/blog/caretta/)
* [Patrick Rogers, August 2023, BPFAgent: eBPF for Monitoring at DoorDash](https://doordash.engineering/2023/08/15/bpfagent-ebpf-for-monitoring-at-doordash/)
* https://github.com/iovisor/bcc
* https://github.com/iovisor/bpftrace

## Projects

* [kubearmor/KubeArmor](https://github.com/kubearmor/KubeArmor/)
* [kubescape/kubescape](https://github.com/kubescape/kubescape/)
* [facebookincubator/dns](https://github.com/facebookincubator/dns/)
* [groundcover-com/caretta](https://github.com/groundcover-com/caretta/)
* [deepflowio/deepflow](https://github.com/deepflowio/deepflow/)
* [grafana/beyla](https://github.com/grafana/beyla/)
* [inspektor-gadget/inspektor-gadget](https://github.com/inspektor-gadget/inspektor-gadget/)
* [slimtoolkit/slim](https://github.com/slimtoolkit/slim/)
* [parca-dev/parca](https://github.com/parca-dev/parca)
* [loxilb-io/loxilb](https://github.com/loxilb-io/loxilb/)
* [oracle/bpftune](https://github.com/oracle/bpftune/)
* [cilium/tetragon](https://github.com/cilium/tetragon/)
* [ddosify/alaz](https://github.com/ddosify/alaz)
* [Anuj Srivastava, October 2022, Skyfall: eBPF agent for infrastructure observability](https://engineering.linkedin.com/blog/2022/skyfall--ebpf-agent-for-infrastructure-observability)

## Security

* [Leonardo Di Donato, Elastic & KP Singh, Google, October 2021, LSM BPF Change Everything](https://www.youtube.com/watch?v=l8jZ-8uLdVU)
* [Rex Guo, Junyuan Zeng, DEF CON 29, Phantom Attack: Evading System Call Monitoring](https://www.youtube.com/watch?v=yaAdM8pWKG8)
* [Michael Zandi, December 2021, BlackBerry Blog, Reverse Engineering Ebpfkit Rootkit With BlackBerry's Enhanced IDA Processor Tool](https://blogs.blackberry.com/en/2021/12/reverse-engineering-ebpfkit-rootkit-with-blackberrys-free-ida-processor-tool)

## Companies

| Name                                        | Description                                                                                                          |
| ------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| [Groundcover](https://www.groundcover.com/) | Observability, for the Cloud. Monitor everything you run in K8s without compromising on cost, granularity, or scale. |
| [Kong](https://konghq.com/)                 | [Kong](https://github.com/Kong/kong/)                                                                                |
| [Grafana](https://grafana.com/)             | [Beyla](https://github.com/grafana/beyla/)                                                                           |
| [DataDog](https://www.datadoghq.com/)       | [DataDog Agent](https://github.com/DataDog/datadog-agent/)                                                           |
| [Isovalent](https://isovalent.com/)         | [Tetragon](https://github.com/cilium/tetragon/), [cilium/ebpf](https://github.com/cilium/ebpf)                       |
| [Cycode](https://cycode.com/)               | [CI/MON](https://cycode.com/cimon-build-hardening/)                                                                  |
| [Aqua Security](https://www.aquasec.com/)   | [Tracee](https://github.com/aquasecurity/tracee)                                                                     |

## People

* Alexei Starovoitov (Facebook, creator of extended BPF)
* Daniel Borkmann (Isovalent)
* David S. Miller (Red Hat)
* Jakub Kicinski (Netronome)
* Yonghong Song (Facebook)
* Martin KaFai Lau (Facebook)
* John Fastabend (Isovalent; formerly Intel)
* Quentin Monnet (Netronome)
* Jesper Dangaard Brouer (Red Hat)
* Andrey Ignatov (Facebook)
* Stanislav Fomichev (Google)
* Brenden Blanco (VMware; formerly PLUMgrid)
* Sasha Goldshtein (Google; formerly SELA)
* Paul Chaignon (Orange)
* Vicent Martí (github)
* Mark Drayton (Facebook)
* Allan McAleavy (Sky)
* Gary Ching-Pang Lin (SUSE)
* Alastair Robertson (Yellowbrick Data; formerly G-Research, Cisco)
* Matheus Marchini (Netflix; formerly Shtima)
* Willian Gasper (Shtima)
* Dale Hamel (Shopify)
* Augusto Mecking Caringi (Red Hat)
* Dan Xu (Facebook)
* Tobias Waldekranz
* Chandler Carruth (Google)
* Mathieu Desnoyers (EfficiOS)
* Jonathan Corbet (LWN.net)
* Jake Edge (LWN.net)
* Eric Dumazet (BPF JIT)
* Will Drewry (BPF filters for seccomp syscall policies)

## Uncategorized

* [Thomas Graf, October 2016, Docker Distributed System Summit, Cilium - BPF & XDP for containers](https://www.youtube.com/watch?v=TnJF7ht3ZYc)
* [Thomas Graf, April 2017, Cilium: Network and Application Security with BPF and XDP](https://www.youtube.com/watch?v=ilKlmTDdFgk)
* [Daniel Borkmann, September 2020, eBPF and Kubernetes: Little Helper Minions for Scaling Microservices](https://www.youtube.com/watch?v=99jUcLt3rSk)
* [Jonathan Corbet, April 2021, LWN.net, Toward signed BPF programs](https://lwn.net/Articles/853489/)
* [July 2022, Introduction to CAP_BPF](https://mdaverde.com/posts/cap-bpf/)
* [Brendan Gregg, November 2017, USENiX, LISA17 - Linux Container Performance Analysis](https://www.youtube.com/watch?v=NYLXZ58EboM)
* [Wenbo Zhang, December 2020, PingCAP Engineering, Trace Linux System Calls with Least Impact on Performance in Production](https://www.pingcap.com/blog/how-to-trace-linux-system-calls-in-production-with-minimal-impact-on-performance/)