// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.Team
{
    public static class GetTeam
    {
        /// <summary>
        /// Use this data source to get information about a Heroku Team.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl-terraform
        /// data "heroku_team" "my_heroku_team" {
        ///   name = "name_of_my_heroku_team"
        /// }
        /// 
        /// output "heroku_team_data_basic" {
        ///   value = [
        ///     "Heroku team",
        ///     "id: ${data.heroku_team.my_heroku_team.id}",
        ///     "default: ${data.heroku_team.my_heroku_team.default}",
        ///     "membership_limit: ${data.heroku_team.my_heroku_team.membership_limit}",
        ///     "provisioned_licenses: ${data.heroku_team.my_heroku_team.provisioned_licenses}",
        ///     "type: ${data.heroku_team.my_heroku_team.type}",
        ///   ]
        /// }
        /// ```
        /// </summary>
        public static Task<GetTeamResult> InvokeAsync(GetTeamArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTeamResult>("heroku:team/getTeam:getTeam", args ?? new GetTeamArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about a Heroku Team.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl-terraform
        /// data "heroku_team" "my_heroku_team" {
        ///   name = "name_of_my_heroku_team"
        /// }
        /// 
        /// output "heroku_team_data_basic" {
        ///   value = [
        ///     "Heroku team",
        ///     "id: ${data.heroku_team.my_heroku_team.id}",
        ///     "default: ${data.heroku_team.my_heroku_team.default}",
        ///     "membership_limit: ${data.heroku_team.my_heroku_team.membership_limit}",
        ///     "provisioned_licenses: ${data.heroku_team.my_heroku_team.provisioned_licenses}",
        ///     "type: ${data.heroku_team.my_heroku_team.type}",
        ///   ]
        /// }
        /// ```
        /// </summary>
        public static Output<GetTeamResult> Invoke(GetTeamInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTeamResult>("heroku:team/getTeam:getTeam", args ?? new GetTeamInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about a Heroku Team.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl-terraform
        /// data "heroku_team" "my_heroku_team" {
        ///   name = "name_of_my_heroku_team"
        /// }
        /// 
        /// output "heroku_team_data_basic" {
        ///   value = [
        ///     "Heroku team",
        ///     "id: ${data.heroku_team.my_heroku_team.id}",
        ///     "default: ${data.heroku_team.my_heroku_team.default}",
        ///     "membership_limit: ${data.heroku_team.my_heroku_team.membership_limit}",
        ///     "provisioned_licenses: ${data.heroku_team.my_heroku_team.provisioned_licenses}",
        ///     "type: ${data.heroku_team.my_heroku_team.type}",
        ///   ]
        /// }
        /// ```
        /// </summary>
        public static Output<GetTeamResult> Invoke(GetTeamInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetTeamResult>("heroku:team/getTeam:getTeam", args ?? new GetTeamInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTeamArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The team name
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetTeamArgs()
        {
        }
        public static new GetTeamArgs Empty => new GetTeamArgs();
    }

    public sealed class GetTeamInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The team name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetTeamInvokeArgs()
        {
        }
        public static new GetTeamInvokeArgs Empty => new GetTeamInvokeArgs();
    }


    [OutputType]
    public sealed class GetTeamResult
    {
        /// <summary>
        /// Whether to use this team when none is specified
        /// </summary>
        public readonly bool Default;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Upper limit of members allowed in a team
        /// </summary>
        public readonly int MembershipLimit;
        public readonly string Name;
        /// <summary>
        /// Whether the team is provisioned licenses by Salesforce
        /// </summary>
        public readonly bool ProvisionedLicenses;
        /// <summary>
        /// type of team Will likely be either "enterprise" or "team"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetTeamResult(
            bool @default,

            string id,

            int membershipLimit,

            string name,

            bool provisionedLicenses,

            string type)
        {
            Default = @default;
            Id = id;
            MembershipLimit = membershipLimit;
            Name = name;
            ProvisionedLicenses = provisionedLicenses;
            Type = type;
        }
    }
}
