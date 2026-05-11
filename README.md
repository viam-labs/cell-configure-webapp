# cell-configure-webapp

Configure-time webapp for a palletizing/picking workcell. Lets an SI or
integrator view, edit, and visualize the cell's `pallet`, `pick-station`,
and `pack-sequencer` in 3D before the operator app runs cycles against it.

This is an apps-only module — no Go-side components or services. The
webapp talks to whatever sibling modules are configured on the same
machine:

- `viam:workcell-components` for the pallet + pick-station components
- `viam:pack-sequencer` for the pack-order service

## Application

| Name | Type | Entrypoint |
|---|---|---|
| `cell-configure-webapp` | `single_machine` | `apps/configure/index.html` |

Launched from the machine's **Apps** tab in the Viam app.

## Build

```
make module.tar.gz
viam module upload --version <X.Y.Z-rcN> --platform linux/amd64 module.tar.gz
```
