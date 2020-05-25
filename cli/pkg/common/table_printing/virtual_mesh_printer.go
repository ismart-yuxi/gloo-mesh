package table_printing

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/solo-io/service-mesh-hub/cli/pkg/common/table_printing/internal"
	zephyr_core_types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	zephyr_networking "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1"
	zephyr_networking_types "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1/types"
	"github.com/solo-io/service-mesh-hub/pkg/security/certgen"
)

func NewVirtualMeshPrinter(tableBuilder TableBuilder) VirtualMeshPrinter {
	return &virtualMeshPrinter{
		tableBuilder: tableBuilder,
	}
}

type virtualMeshPrinter struct {
	tableBuilder TableBuilder
}

func (m *virtualMeshPrinter) Print(
	out io.Writer,
	virtualMeshes []*zephyr_networking.VirtualMesh,
) error {
	if len(virtualMeshes) == 0 {
		return nil
	}

	table := m.tableBuilder(out)

	commonHeaderRows := []string{
		"Metadata",
		"Meshes",
		"Config",
		"Status",
	}
	var preFilteredRows [][]string
	for _, virtualMesh := range virtualMeshes {
		var newRow []string
		// Append common metadata
		newRow = append(newRow, m.buildMetadataCell(virtualMesh))

		// Append meshes
		newRow = append(newRow, m.buildMeshesCell(virtualMesh.Spec.GetMeshes()))

		// Append federation data
		cell, err := m.buildConfigCell(virtualMesh.Spec)
		if err != nil {
			return err
		}
		newRow = append(newRow, cell)

		// Append status data
		newRow = append(newRow, m.buildStatusCell(virtualMesh.Status))

		preFilteredRows = append(preFilteredRows, newRow)
	}

	filteredHeaders, filteredRows := internal.FilterEmptyColumns(commonHeaderRows, preFilteredRows)

	table.SetHeader(filteredHeaders)
	table.AppendBulk(filteredRows)
	table.Render()
	return nil
}

func (m *virtualMeshPrinter) buildMetadataCell(virtualMesh *zephyr_networking.VirtualMesh) string {
	var items []string

	items = append(items, fmt.Sprintf("Name: %s", virtualMesh.GetName()))
	items = append(items, fmt.Sprintf("Namespace: %s", virtualMesh.GetNamespace()))
	items = append(items, fmt.Sprintf("Display Name: %s", virtualMesh.Spec.GetDisplayName()))

	return strings.Join(items, "\n")
}

func (m *virtualMeshPrinter) buildMeshesCell(meshList []*zephyr_core_types.ResourceRef) string {
	if len(meshList) == 0 {
		return ""
	}
	var items []string

	for _, mesh := range meshList {
		items = append(items, fmt.Sprintf("- %s", mesh.GetName()))
	}

	return strings.Join(items, "\n")
}

func (m *virtualMeshPrinter) buildConfigCell(
	spec zephyr_networking_types.VirtualMeshSpec,
) (string, error) {
	var items []string

	switch spec.GetTrustModel().(type) {
	case *zephyr_networking_types.VirtualMeshSpec_Limited:
		items = append(items, "Trust Mode: Limited")
	default:
		// Defaults to shared
		items = append(items, "Trust Mode: Shared")
	}

	switch certType := spec.GetCertificateAuthority().GetType().(type) {
	case *zephyr_networking_types.VirtualMeshSpec_CertificateAuthority_Builtin_:
		items = append(items,
			"\nCertificate Authority:",
			"  Type: Self Signed",
		)
		orgName := certgen.DefaultOrgName
		if certType.Builtin.GetOrgName() != "" {
			orgName = certType.Builtin.GetOrgName()
		}
		items = append(items, fmt.Sprintf("  Org Name: %s", orgName))
		ttlDays := strconv.Itoa(certgen.DefaultRootCertTTLDays)
		if certType.Builtin.GetTtlDays() != 0 {
			ttlDays = string(certType.Builtin.GetTtlDays())
		}
		items = append(items, fmt.Sprintf("  TTL: %s days", ttlDays))
		keySize := certgen.DefaultRootCertRsaKeySize
		if certType.Builtin.GetRsaKeySizeBytes() != 0 {
			keySize = int(certType.Builtin.GetRsaKeySizeBytes())
		}
		items = append(items, fmt.Sprintf("  Key Size: %d", keySize))
	case *zephyr_networking_types.VirtualMeshSpec_CertificateAuthority_Provided_:
		items = append(items,
			"\nCertificate Authority:",
			"  Type: Provided",
			fmt.Sprintf("  Secret Name: %s", certType.Provided.GetCertificate().GetName()),
			fmt.Sprintf("  Secret Namespace: %s", certType.Provided.GetCertificate().GetNamespace()),
		)
	default:
		items = append(items,
			"\nCertificate Authority:",
			"  Type: Self Signed",
			fmt.Sprintf("  Org Name: %s", certgen.DefaultOrgName),
			fmt.Sprintf("  TTL: %d days", certgen.DefaultRootCertTTLDays),
			fmt.Sprintf("  Key Size: %d", certgen.DefaultRootCertRsaKeySize),
		)
	}

	items = append(items, fmt.Sprintf("\nFederation Mode: %s", spec.GetFederation().GetMode().String()))

	accessControlEnforcement := spec.GetEnforceAccessControl().String()
	items = append(items, fmt.Sprintf("\nAccess Control Enforcement: %s", accessControlEnforcement))

	return strings.Join(items, "\n"), nil
}

func (m *virtualMeshPrinter) buildStatusCell(status zephyr_networking_types.VirtualMeshStatus) string {
	var items []string

	if federationCell := BuildSimpleStatusCell(
		status.GetFederationStatus(),
		"Federation",
	); federationCell != "" {
		items = append(items, federationCell)
	}

	if certCell := BuildSimpleStatusCell(status.GetCertificateStatus(), "Certificate"); certCell != "" {
		items = append(items, certCell)
	}
	if configCell := BuildSimpleStatusCell(status.GetConfigStatus(), "Config"); configCell != "" {
		items = append(items, configCell)
	}
	if acpCell := BuildSimpleStatusCell(
		status.GetAccessControlEnforcementStatus(),
		"Access Control Enforcement",
	); acpCell != "" {
		items = append(items, acpCell)
	}

	return strings.Join(items, "\n\n")
}
