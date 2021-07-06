#!/bin/bash
set -x
# $ curl -O -L https://releases-art-rhcos.svc.ci.openshift.org/art/storage/releases/rhcos-4.6/46.82.202007051540-0/x86_64/rhcos-46.82.202007051540-0-qemu.x86_64.qcow2.gz
# $ mv rhcos-46.82.202007051540-0-qemu.x86_64.qcow2.gz /tmp
# $ sudo gunzip /tmp/rhcos-46.82.202007051540-0-qemu.x86_64.qcow2.gz

VM_NAME="$1"
DOMAIN="example.com"
CLUSTER_NAME="cluster1"
IGNITION_CONFIG="/var/lib/libvirt/images/aio.ign"
cp "${VM_NAME}/aio.ign" "${IGNITION_CONFIG}"
chown qemu:qemu "${IGNITION_CONFIG}"
restorecon "${IGNITION_CONFIG}"

RHCOS_IMAGE="/tmp/rhcos-46.82.202007051540-0-qemu.x86_64.qcow2"

INSTALL_ISO="/home/qemu-virt/discovery_image_${CLUSTER_NAME}.iso"
OS_VARIANT="rhel8.1"
VCPUS="8"
# RAM_MB="16384"
RAM_MB="24576"
DISK_GB="60"
DISK="/home/qemu-virt/${VM_NAME}.cow"
rm -f "${DISK}"

LASTOCT=$(echo ${VM_NAME} | md5sum | cut -c 1-2)
MAC="52:54:00:ee:42:$LASTOCT"

NETWORK=test-net

virsh net-dhcp-leases $NETWORK | grep ${MAC}
    #if [ $? = 0 ]; then
virsh net-update $NETWORK delete ip-dhcp-host "<host mac='${MAC}'/>" --live 2>&1 > /dev/null
    #fi

ip=10
network=$( virsh net-dumpxml $NETWORK | grep range | awk -F \' '{print $2}' | awk -F . '{print $1"."$2"."$3}' )
rc=1
while [ 1 = 1 ]; do
    virsh net-update $NETWORK add ip-dhcp-host "<host mac='${MAC}' name='${VM_NAME}' ip='${network}.${ip}'/>" --live
    rc=$?
    if [ $rc = 0 ]; then
        break
    elif [ $ip -gt 254 ]; then
	echo "No more IPs left"
        exit 1
    else
        ip=$(($ip + 1))
    fi 
done

DNS=$(grep baseDomain install-config-${VM_NAME}.yaml | cut -d: -f2 | tr -d ' \t')

echo "" > /etc/NetworkManager/dnsmasq.d/${VM_NAME}.conf
if [ ${VM_NAME} = ${CLUSTER_NAME} ]; then
        echo "address=/.${CLUSTER_NAME}.${DOMAIN}/${network}.${ip}" >> /etc/NetworkManager/dnsmasq.d/${VM_NAME}.conf
        # echo "address=/.${VM_NAME}.${DNS}/${network}.${ip}" > /etc/NetworkManager/dnsmasq.d/${VM_NAME}.conf
fi
echo "address=/.${VM_NAME}/${network}.${ip}" >> /etc/NetworkManager/dnsmasq.d/${VM_NAME}.conf
sudo systemctl reload NetworkManager.service

echo "Installing ${VM_NAME} ( $MAC ) @ ${network}.${ip}"

virt-install \
    --connect qemu:///system \
    -n "${VM_NAME}" \
    -r "${RAM_MB}" \
    --vcpus "${VCPUS}" \
    --os-variant="${OS_VARIANT}" \
    --import \
    --network=network:test-net,mac="$MAC" \
    --graphics=none \
    --disk "path=${DISK},size=${DISK_GB}" \
    --cdrom "${INSTALL_ISO}" \
    # --unattended
    # --noautoconsole \
    # --disk "path=${DISK},size=${DISK_GB},backing_store=${RHCOS_IMAGE}" \
    # --qemu-commandline="-fw_cfg name=opt/com.coreos/config,file=${IGNITION_CONFIG}"

echo "Waiting 5 minutes for the node to bootstrap"
sleep 300
