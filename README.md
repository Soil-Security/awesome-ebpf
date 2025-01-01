# Awesome Berkeley Packet Filter (BPF)

A very opinionated set of awesome resources that I read or bumped into while leraning eBPF.

## Fundamentals

* [McCanne S., Jacobson V., Dec 1992, The BSD Packet Filter: A New Architecture for User-level Packet Capture](https://www.tcpdump.org/papers/bpf-usenix93.pdf)
* [Corbet J., June 2021, Spectre revisits BPF, LWN.net](https://lwn.net/Articles/860597/)
* [Corbet J., Apr 2011, A JIT for packet filters, LWN.net](https://lwn.net/Articles/437981/)
* [Corbet J., Jan 2012, Yet another new approach to seccomp, LWN.net](https://lwn.net/Articles/475043/)
* [Starovoitov A., Sept 2013, extended BPF, LKML.org](https://lkml.org/lkml/2013/9/30/627)
* [Starovoitov A., Aug 2022, Kernel Recipes 2022 - The untold story of BPF](https://www.youtube.com/watch?v=DAvZH13725I)
* [Fastabend J., Dec 2022, Isovalent, How the Hive Came to Bee - The History of eBPF](https://www.youtube.com/watch?v=sYxUOKi3Q00)
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
* [eBPF: What is it, Best Practices, and Use Cases, Groundcover](https://www.groundcover.com/ebpf)
* [Short C., Aug 2022, Intro to eBPF, chrisshort.net](https://chrisshort.net/intro-to-ebpf/)
* [Mulligan B., Borkmann D., Apr 2023, The Silent Platform Revolution: How eBPF Is Fundamentally Transforming Cloud-Native Platforms, InfoQ](https://www.infoq.com/articles/ebpf-cloud-native-platforms/)
* [Podobnik T. J., Sep 2024, Loops and Iterators in eBPF, eBPFChirp](https://ebpfchirp.substack.com/p/loops-and-iterators-in-ebpf)
* [Isovalent eBPF Docs](https://docs.ebpf.io/)

## Advanced

* [Rice L., Fastabend J., Sept 2024, eBPF: Yes, it’s Turing Complete!, Isovalent Blog](https://isovalent.com/blog/post/ebpf-yes-its-turing-complete/)
* [Edge J., Sept 2019, Kernel runtime security instrumentation, LWN.net](https://lwn.net/Articles/798157/)
* [Fleming M., May 2018, Using user-space tracepoints with BPF, LWN.net](https://lwn.net/Articles/753601/)
* [Fleming M., Dec 2017, A thorough introduction to eBPF, LWN.net](https://lwn.net/Articles/740157/)
* [Corbet J., Dec 2018, Bounded loops in BPF programs, LWN.net](https://lwn.net/Articles/773605/)
* [Corbet J., May 2014, BPF: the universal in-kernel virtual machine, LWN.net](https://lwn.net/Articles/599755/)
* [Corbet J., Jul 2014, Extending extended BPF, LWN.net](https://lwn.net/Articles/603983/)
* [Corbet J., Sept 2014, The BPF system call API, version 14, LWN.net](https://lwn.net/Articles/612878/)
* [Goswami S., Apr 2005, An introduction to KProbes, LWN.net](https://lwn.net/Articles/132196/)
* [Nurkholish Halim, Sept 2022, Medium, A Deep Dive into eBPF: Writing an Efficient DNS Monitoring](https://medium.com/@nurkholish.halim/a-deep-dive-into-ebpf-writing-an-efficient-dns-monitoring-2c9dea92abdf)
* [Adrian Ratiu, Apr 2019, Collabora, An eBPF overview, part 1: Introduction](https://www.collabora.com/news-and-blog/blog/2019/04/05/an-ebpf-overview-part-1-introduction/)
* [Adrian Ratiu, Apr 2019, Collabora, An eBPF overview, part 2: Machine & bytecode](https://www.collabora.com/news-and-blog/blog/2019/04/15/an-ebpf-overview-part-2-machine-and-bytecode/)
* [Securing Linux with a Faster and Scalable IPtables](https://ccronline.sigcomm.org/wp-content/uploads/2019/07/acmdl19-304.pdf)
* [Luuk Hendriks, Jul 2020, RIPE Labs, Journeying into XDP: Part 0](https://labs.ripe.net/author/luuk_hendriks/journeying-into-xdp-part-0/)
* [Carpay T., Oct 2020, RIPE Labs, Journeying into XDP Part 1: Augmenting DNS](https://labs.ripe.net/author/tom_carpay/journeying-into-xdp-part-1-augmenting-dns/)
* [Luuk Hendriks, Jul 2020, RIPE Labs, Journeying into XDP Part 2: XDPerimenting with DNS Telemetry](https://labs.ripe.net/author/luuk_hendriks/journeying-into-xdp-part-2-xdperimenting-with-dns-telemetry/)
* [Sanjeev Rampal, Donald Hunter, May 2023, eBPF 201: Supercharging Your eBPF Dev Process for Cloud Native Apps](https://www.youtube.com/watch?v=9vUJhbqf4eo)

## XDP

* [Oct 2018, Load XDP programs using the ip (iproute2) command](https://fntlnz.wtf/post/xdp-ip-iproute/)
* [Oct 2022, Writing an eBPF/XDP load-balancer in Rust, Kong Blog](https://konghq.com/blog/engineering/writing-an-ebpf-xdp-load-balancer-in-rust)

## Computer Science

* [Wikipedia, Control-flow graph](https://en.wikipedia.org/wiki/Control-flow_graph)
* [Yunhe Shi, David Gregg, Andrew Beatty, M. Anton Ertl, Virtual Machine Showdown: Stack Versus Registers](https://www.usenix.org/legacy/events/vee05/full_papers/p153-yunhe.pdf)
* [Hongjiu Lu, May 1995, ELF: From The Programmer’s Perspective](https://www.linux.co.cr/free-unix-os/review/acrobat/950517.pdf)
* [June 2019, Simple and Precise Static Analysis of Untrusted Linux Kernel Extensions](https://dl.acm.org/doi/pdf/10.1145/3314221.3314590)
* [The eXpress data path: fast programmable packet processing in the operating system kernel](https://dl.acm.org/doi/10.1145/3281411.3281443)

## Books

* [Rice L., Mar 2023, O'Reilly, Learning eBPF](https://www.oreilly.com/library/view/learning-ebpf/9781098135119/)
* [Calavera D., Fontana L., Sept 2019, Linux Observability with BPF](https://www.oreilly.com/library/view/linux-observability-with/9781492050193/)
* [Gregg B., Dec 2019, Addison-Wesley Professional, BPF Performance Tools](https://www.oreilly.com/library/view/bpf-performance-tools/9780136588870/)
* [Gregg B., Nov 2020, Addison-Wesley, Systems Performance: Enterprise and the Cloud](https://www.brendangregg.com/blog/2020-07-15/systems-performance-2nd-edition.html)
* [Degioanni L., Grasso L., Aug 2022, O'Reilly, Practical Cloud Native Security with Falco](https://www.oreilly.com/library/view/practical-cloud-native/9781098118563/)
* [Salazar J., Reka Ivanko N., Apr 2022, O'Reilly, Security Observability with eBPF](https://www.oreilly.com/library/view/security-observability-with/9781492096719/)
* [Rice L., Apr 2022, O'Reilly, What is eBPF?](https://www.oreilly.com/library/view/what-is-ebpf/9781492097266/)

## Compilers

* [Monnet Q., Apr 2020, eBPF assembly with LLVM, Whirl Offload](https://qmonnet.github.io/whirl-offload/2020/04/12/llvm-ebpf-asm/)
* [Corbet J., Sept 2019, Compiling to BPF with GCC, LWN.net](https://lwn.net/Articles/800606/)
* [Edge J., Sept 2020, BPF in GCC, LWN.net](https://lwn.net/Articles/831402/)
* [Jiong Wang, Sept 2018, Netronome, Demystify eBPF JIT Compiler Webinar](https://www.youtube.com/watch?v=KDr033U2fF4)

## Code Examples

* [Soil-Security/bpf-c-template](https://github.com/Soil-Security/bpf-c-template/)
* [Soil-Security/bpf-go-template](https://github.com/Soil-Security/bpf-go-template/)
* [Soil-Security/bpf-xdp-go-template](https://github.com/Soil-Security/bpf-xdp-go-template/)
* [Soil-Security/bpf-go-daemonset-template](https://github.com/Soil-Security/bpf-go-daemonset-template/)
* [Soil-Security/bpf-go-socket](https://github.com/Soil-Security/bpf-go-socket/)
* [:lock: danielpacak/bpf-xdp-rust-template](https://github.com/danielpacak/bpf-xdp-rust-template/)
* [:lock: danielpacak/bpf-tc-go-template](https://github.com/danielpacak/bpf-tc-go-template/)
* [:lock: danielpacak/bpf-windows-template](https://github.com/danielpacak/bpf-windows-template/)
* [libbpf-bootstrap: demo BPF applications](https://github.com/libbpf/libbpf-bootstrap/)
* [isovalent/game-of-life](https://github.com/isovalent/game-of-life)
* [David Calavera, Lorenzo Fontana, BPF Workshop](https://github.com/bpftools/bpf-workshop/)
* [xdp-project/xdp-tutorial](https://github.com/xdp-project/xdp-tutorial/)
* [NLnetLabs/XDPeriments](https://github.com/NLnetLabs/XDPeriments/)
* https://eunomia.dev/

## Tools and Frameworks

* [Monnet Q., Sept 2021, Features of bpftool: the thread of tips and examples to work with eBPF objects, Whirl Offload](https://qmonnet.github.io/whirl-offload/2021/09/23/bpftool-features-thread/)
* [Kerrisk M., Feb 2020, Linux Security and Isolation APIs Seccomp](https://man7.org/training/download/secisol_seccomp_slides.pdf)
* [edgebitio/edgebit-agent](https://github.com/edgebitio/edgebit-agent)
* [:lock: danielpacak/lazybpftool](https://github.com/danielpacak/lazybpftool)
* [bpfd-dev/bpfd](https://github.com/bpfd-dev/bpfd)
* [rust-bpf/rust-bcc](https://github.com/rust-bpf/rust-bcc)
* [Oct 2021, eCHO episode 25: eBPF, Rust and Aya](https://www.youtube.com/watch?v=TQ0ou-eFLAk)
* [Monnet Q., Feb 2020, Tools and mechanisms to debug BPF programs](https://archive.fosdem.org/2020/schedule/event/debugging_bpf/)
* [Udi Rot, Jan 2023, Navigate your way to production bliss with Caretta](https://www.groundcover.com/blog/caretta/)
* [Rogers P., Aug 2023, BPFAgent: eBPF for Monitoring at DoorDash](https://doordash.engineering/2023/08/15/bpfagent-ebpf-for-monitoring-at-doordash/)
* https://github.com/iovisor/bcc
* https://github.com/iovisor/bpftrace
* https://github.com/google/buzzer

## Projects

* [Gerring S., Jul 2024, eBPF Network Vershitifier, Scott's Ramblings](https://blog.scottgerring.com/ebpf-network-vershitifier/)
* [open-telemetry/opentelemetry-ebpf-profiler](https://github.com/open-telemetry/opentelemetry-ebpf-profiler)
* [coroot/coroot](https://github.com/coroot/coroot)
* [microsoft/retina](https://github.com/microsoft/retina)
* [kubearmor/KubeArmor](https://github.com/kubearmor/KubeArmor/)
* [kubescape/kubescape](https://github.com/kubescape/kubescape/)
* [kubeshark/kubeshark](https://github.com/kubeshark/kubeshark)
* [kubeshark/tracer](https://github.com/kubeshark/tracer)
* [facebookincubator/dns](https://github.com/facebookincubator/dns/)
* [groundcover-com/caretta](https://github.com/groundcover-com/caretta/)
* [deepflowio/deepflow](https://github.com/deepflowio/deepflow/)
* [grafana/beyla](https://github.com/grafana/beyla/)
* [inspektor-gadget/inspektor-gadget](https://github.com/inspektor-gadget/inspektor-gadget/)
* [slimtoolkit/slim](https://github.com/slimtoolkit/slim/)
* [parca-dev/parca](https://github.com/parca-dev/parca)
* [loxilb-io/loxilb](https://github.com/loxilb-io/loxilb/)
* [oracle/bpftune](https://github.com/oracle/bpftune/)
* [oracle-samples/bysyscall](https://github.com/oracle-samples/bysyscall)
* [cilium/tetragon](https://github.com/cilium/tetragon/)
* [ddosify/alaz](https://github.com/ddosify/alaz)
* [keisku/gmon](https://github.com/keisku/gmon)
* [amiremohamadi/bpfsnake](https://github.com/amiremohamadi/bpfsnake)
* [alegrey91/harpoon](https://github.com/alegrey91/harpoon)
* [furkanonder/DnsTrace](https://github.com/furkanonder/DnsTrace)
* [Anuj Srivastava, Oct 2022, Skyfall: eBPF agent for infrastructure observability](https://engineering.linkedin.com/blog/2022/skyfall--ebpf-agent-for-infrastructure-observability)
* [Sumera Priyadarsini, Jul 2022, Polar Signals Blog, Introduction to Parca - Part 1](https://www.polarsignals.com/blog/posts/2022/07/12/introducing-parca-sequel/)
* [citronneur/pamspy](https://github.com/citronneur/pamspy)

## Kubernetes

* [Podobnik T. J., Oct 2024, Securing Kubernetes Workloads using LSM-BPF, eBPFChirp](https://ebpfchirp.substack.com/p/securing-kubernetes-workloads-using)
* [Calavera D., Dec 2018, Spy on your Kubernetes cluster with BPF, Medium](https://medium.com/@calavera/spy-on-your-kubernetes-cluster-with-bpf-b09032bd1cdc)

## Security

* [Dinaburg A., Sept 2023, Pitfalls of relying on eBPF for security monitoring (and some solutions), Trail of Bits Blog](https://blog.trailofbits.com/2023/09/25/pitfalls-of-relying-on-ebpf-for-security-monitoring-and-some-solutions/)
* [Jun 2022, Bypassing eBPF-based Security Enforcement Tools, Form3 Blog](https://www.form3.tech/blog/engineering/bypassing-ebpf-tools)
* [Leonardo Di Donato, Elastic & KP Singh, Google, Oct 2021, LSM BPF Change Everything](https://www.youtube.com/watch?v=l8jZ-8uLdVU)
* [Rex Guo, Junyuan Zeng, DEF CON 29, Phantom Attack: Evading System Call Monitoring](https://www.youtube.com/watch?v=yaAdM8pWKG8)
* [Zandi M., Dec 2021, BlackBerry Blog, Reverse Engineering Ebpfkit Rootkit With BlackBerry's Enhanced IDA Processor Tool](https://blogs.blackberry.com/en/2021/12/reverse-engineering-ebpfkit-rootkit-with-blackberrys-free-ida-processor-tool)
* [Spyderbat, Aug2023, How eBPF Can Help Identify Container Escapes](https://www.spyderbat.com/blog/how-ebpf-can-help-identify-container-escapes)
* [Spyderbat, Aug 2023, Using eBPF to Resolve GuardDuty DNS Alerts](https://www.spyderbat.com/blog/using-ebpf-to-resolve-guardduty-dns-alerts)
* [KP Singh, Google, Nov 2020, Security Auditing and Enforcement using eBPF, eBPF Summit 2020](https://www.youtube.com/watch?v=XFJw37Vwzcc)
* [Juan José López Jaimez, Inge M., Aug 2024, A deep dive into CVE-2023-2163: How we found and fixed an eBPF Linux Kernel Vulnerability, Google Bug Hunters](https://bughunters.google.com/blog/6303226026131456/a-deep-dive-into-cve-2023-2163-how-we-found-and-fixed-an-ebpf-linux-kernel-vulnerability)
* [Kondah M., Jun 2024, Profiling Libraries With eBPF: Detecting Zero-Day Exploits and Backdoors, Deep Kondah](https://www.deep-kondah.com/ebpf-library-profiling/)
* [The Risks of Using eBPF for Security, Vali Cyber](https://valicyber.com/resources/the-risks-of-using-ebpf-for-security/)
* [Gadient A., Nov 2024, EBPF-Based Security Solutions: Exploring Weaknesses And Mitigation Techniques., Forbes](https://www.forbes.com/councils/forbestechcouncil/2023/11/15/ebpf-based-security-solutions-exploring-weaknesses-and-mitigation-techniques/)

## Windows

* [Use eBPF-based sensor for Microsoft Defender for Endpoint on Linux, Microsoft Defender Documentation](https://learn.microsoft.com/en-us/defender-endpoint/linux-support-ebpf)

## Companies

| Name                                        | Description                                                                                                          |
| ------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| [Spyderbat](https://www.spyderbat.com/)     |                                                                                                                      |
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

## Stack Overflow

* [full path for open / openat relative filenames](https://stackoverflow.com/questions/68240576/full-path-for-open-openat-relative-filenames)

## Uncategorized

* [Graf T., Oct 2016, Docker Distributed System Summit, Cilium - BPF & XDP for containers](https://www.youtube.com/watch?v=TnJF7ht3ZYc)
* [Graf T., Apr 2017, Cilium: Network and Application Security with BPF and XDP](https://www.youtube.com/watch?v=ilKlmTDdFgk)
* [Borkmann D., Sept 2020, eBPF and Kubernetes: Little Helper Minions for Scaling Microservices](https://www.youtube.com/watch?v=99jUcLt3rSk)
* [Corbet J., Apr 2021, LWN.net, Toward signed BPF programs](https://lwn.net/Articles/853489/)
* [Landaverde M., Jul 2022, Introduction to CAP_BPF](https://mdaverde.com/posts/cap-bpf/)
* [Gregg B., Nov 2017, USENiX, LISA17 - Linux Container Performance Analysis](https://www.youtube.com/watch?v=NYLXZ58EboM)
* [Wenbo Zhang, Dec 2020, Trace Linux System Calls with Least Impact on Performance in Production, PingCAP Engineering](https://www.pingcap.com/blog/how-to-trace-linux-system-calls-in-production-with-minimal-impact-on-performance/)
* [Wenbo Zhang, Dec 2020, Tips and Tricks for Writing Linux BPF Applications with libbpf, PingCAP Engineering](https://www.pingcap.com/blog/tips-and-tricks-for-writing-linux-bpf-applications-with-libbpf/)
* [Majkowski M., Mar 2018, eBPF, Sockets, Hop Distance and manually writing eBPF assembly](https://blog.cloudflare.com/epbf_sockets_hop_distance/)
* [Hossain R., Jul 2023, [Tutorial] How eBPF Improves Observability within Kubernetes, Loft Blog](https://loft.sh/blog/tutorial-how-ebpf-improves-observability-within-kubernetes/)

## See Also

* https://github.com/zoidyzoidzoid/awesome-ebpf
